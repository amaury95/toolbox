package gen

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stoewer/go-strcase"
)

func GenerateProto(reader io.Reader, typeName, packageName string, outputPaths ...string) {
	evmABI, err := abi.JSON(reader)
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Imports
	var importEmpty bool

	definitions := getDefinitions(evmABI, &importEmpty)

	eventSubscriptions, eventParams := getEventSubscriptions(evmABI, &importEmpty)

	rpcCalls, rpcParams := getRpcCalls(evmABI, &importEmpty)

	proto := new(bytes.Buffer)

	// Header
	proto.WriteString("syntax = \"proto3\";\n")
	proto.WriteString(fmt.Sprintf("package %s;\n\n", packageName))

	// Imports
	if importEmpty {
		proto.WriteString("import \"google/protobuf/empty.proto\";\n")
	}

	// Definitions
	proto.WriteString(strings.Join(definitions, "\n"))

	// Service
	proto.WriteString("\n\n/* Service */")
	proto.WriteString(fmt.Sprintf("\nservice %sService {\n", typeName))

	// Event Subscriptions and RPC Calls
	proto.WriteString(strings.Join(append(eventSubscriptions, rpcCalls...), "\n"))
	proto.WriteString("\n}\n\n")

	// RPC Params
	proto.WriteString(strings.Join(append(eventParams, rpcParams...), "\n"))

	// Stdout Output
	if len(outputPaths) == 0 || outputPaths[0] == "" {
		fmt.Println(proto.String())
		return
	}

	os.WriteFile(outputPaths[0], proto.Bytes(), 0644)
}

func getEventSubscriptions(evmABI abi.ABI, importEmpty *bool) (eventSubscriptions, eventParams []string) {
	eventSubscriptions = append(eventSubscriptions, "\t/* Event Subscriptions */")
	eventParams = append(eventParams, "/* Event Params */")

	for _, event := range SortMap(evmABI.Events) {
		var request, response string

		if len(indexedInputs(event.Value.Inputs)) == 0 {
			request = "google.protobuf.Empty"
			*importEmpty = true
		} else {
			request = event.Key + "EventRequest"
			eventParams = append(eventParams, definitionMessage(request, indexedInputs(event.Value.Inputs), true))
		}

		if len(event.Value.Inputs) == 0 {
			response = "google.protobuf.Empty"
			*importEmpty = true
		} else {
			response = event.Key + "EventResponse"
			eventParams = append(eventParams, definitionMessage(response, event.Value.Inputs, false))
		}

		eventSubscriptions = append(eventSubscriptions, fmt.Sprintf("\trpc %s(%s) returns (stream %s);", event.Key+`Event`, request, response))
	}
	return
}

func getRpcCalls(evmABI abi.ABI, importEmpty *bool) (rpcCalls, rpcParams []string) {
	actions := []string{"\t/* Actions */"}
	queries := []string{"\t/* Queries */"}
	utils := []string{"\t/* Utils */"}

	params := []string{"/* Params */"}

	for _, method := range SortMap(evmABI.Methods) {
		methodName := strcase.UpperCamelCase(method.Key)
		var request, response string

		if len(method.Value.Inputs) == 0 {
			request = "google.protobuf.Empty"
			*importEmpty = true
		} else {
			request = methodName + "Request"
			params = append(params, definitionMessage(request, method.Value.Inputs, false))
		}

		if len(method.Value.Outputs) == 0 {
			response = "google.protobuf.Empty"
			*importEmpty = true
		} else {
			response = methodName + "Response"
			params = append(params, definitionMessage(response, method.Value.Outputs, false))
		}

		switch method.Value.StateMutability {
		case "pure":
			utils = append(utils, fmt.Sprintf("\trpc %s(%s) returns (%s);", methodName, request, response))
		case "view":
			queries = append(queries, fmt.Sprintf("\trpc %s(%s) returns (%s);", methodName, request, response))
		default:
			actions = append(actions, fmt.Sprintf("\trpc %s(%s) returns (%s);", methodName, request, response))
		}
	}

	rpcCalls = append(rpcCalls, queries...)
	rpcCalls = append(rpcCalls, actions...)
	rpcCalls = append(rpcCalls, utils...)

	rpcParams = append(rpcParams, params...)
	return
}

func getDefinitions(evmABI abi.ABI, _ *bool) (definitions []string) {
	definitions = append(definitions, "\n/* Definitions */")

	found := map[string]string{}

	for _, event := range evmABI.Events {
		for _, input := range event.Inputs {
			for _, def := range walk(&input.Type) {
				found[def.Key] = definitionTuple(*def.Value)
			}
		}
	}

	for _, method := range evmABI.Methods {
		for _, input := range method.Inputs {
			for _, def := range walk(&input.Type) {
				found[def.Key] = definitionTuple(*def.Value)
			}
		}
		for _, output := range method.Outputs {
			for _, def := range walk(&output.Type) {
				found[def.Key] = definitionTuple(*def.Value)
			}
		}
	}

	for _, v := range found {
		definitions = append(definitions, v)
	}
	return
}

func walk(types ...*abi.Type) (results []Tuple[*abi.Type]) {
	for _, t := range types {
		if t.T == abi.SliceTy || t.T == abi.ArrayTy {
			t = t.Elem
		}
		if t.T == abi.TupleTy {
			results = append(results, Tuple[*abi.Type]{Key: t.TupleRawName, Value: t})
			results = append(results, walk(t.TupleElems...)...)
		}
	}
	return
}

func definitionMessage(name string, inputs abi.Arguments, onlyIndexed bool) string {
	message := new(bytes.Buffer)
	message.WriteString(fmt.Sprintf("message %s {\n", name))

	index := 1
	for _, input := range inputs {
		if onlyIndexed && !input.Indexed {
			continue
		}
		message.WriteString(definitionType(input.Name, input.Type, index))
		index++
	}
	message.WriteString("};")
	return message.String()
}

func definitionTuple(t abi.Type) string {
	message := new(bytes.Buffer)
	message.WriteString(fmt.Sprintf("message %s {\n", t.TupleRawName))
	index := 1
	for i, input := range t.TupleElems {
		name := strcase.SnakeCase(t.TupleRawNames[i])
		message.WriteString(definitionType(nameWithoutPrefixAndSuffix(name), *input, index))
		index++
	}
	message.WriteString("};")
	return message.String()
}

func definitionType(name string, t abi.Type, index int) string {
	value := new(bytes.Buffer)

	// Definition name
	defName := strcase.SnakeCase(name)
	if defName == "" && index == 1 {
		defName = "value"
	} else if defName == "" {
		defName = fmt.Sprintf("value%d", index)
	}

	// Handle repeated types
	var repeated bool
	if t.T == abi.ArrayTy || t.T == abi.SliceTy {
		repeated = true
		t = *t.Elem
	}

	value.WriteString("\t")
	if repeated {
		value.WriteString("repeated ")
	}
	value.WriteString(typeOf(t))
	value.WriteString(" ")
	value.WriteString(nameWithoutPrefixAndSuffix(defName))
	value.WriteString(" = ")
	value.WriteString(strconv.Itoa(index))
	value.WriteString(";\n")

	return value.String()
}

func typeOf(t abi.Type) string {
	if t.T == abi.TupleTy {
		return t.TupleRawName
	}
	switch t.T {
	case abi.AddressTy, abi.StringTy:
		return "string"
	case abi.BoolTy:
		return "bool"
	case abi.BytesTy, abi.FixedBytesTy:
		return "bytes"
	case abi.FixedPointTy:
		return "double"
	case abi.IntTy:
		return "int64"
	case abi.UintTy:
		return "uint64"
	default:
		return t.String()
	}
}

func indexedInputs(inputs abi.Arguments) (indexed abi.Arguments) {
	for _, input := range inputs {
		if input.Indexed {
			indexed = append(indexed, input)
		}
	}
	return
}

type Tuple[T any] struct {
	Key   string
	Value T
}

func SortMap[T any](m map[string]T) []Tuple[T] {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	methods := make([]Tuple[T], 0, len(keys))
	for _, k := range keys {
		methods = append(methods, Tuple[T]{Key: k, Value: m[k]})
	}
	return methods
}

func nameWithoutPrefixAndSuffix(name string) string {
	return strings.TrimPrefix(strings.TrimSuffix(name, "_"), "_")
}
