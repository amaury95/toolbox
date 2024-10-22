package command

import (
	"os"

	"github.com/amaury95/toolbox/src/generate"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a password, an ethereum address, etc.",
}

func init() {
	rootCmd.AddCommand(genCmd)
	initMnemonicCmd()
	initEthereumCmd()
	initPasswordCmd()
	initProtoCmd()
}

// Mnemonic

var _mnemonic_name string
var _mnemonic_tags []string
var _mnemonic_encrypt_password string

var mnemonicCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "Generate a Mnemonic",
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateMnemonic(_mnemonic_name, _mnemonic_encrypt_password, _mnemonic_tags...)
	},
}

func initMnemonicCmd() {
	genCmd.AddCommand(mnemonicCmd)
	mnemonicCmd.Flags().StringVarP(&_mnemonic_name, "name", "n", "", "Name for the mnemonic")
	mnemonicCmd.Flags().StringSliceVarP(&_mnemonic_tags, "tag", "t", []string{}, "Tags for the password")
	mnemonicCmd.Flags().StringVarP(&_mnemonic_encrypt_password, "encrypt-password", "e", "", "Password for the encryption")
}

// Ethereum

var _ethereum_tags []string
var _ethereum_encrypt_password string

var ethereumCmd = &cobra.Command{
	Use:   "address",
	Short: "Generate Ethereum private key and address",
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateEthereumKey(_ethereum_encrypt_password, _ethereum_tags...)
	},
}

func initEthereumCmd() {
	genCmd.AddCommand(ethereumCmd)
	ethereumCmd.Flags().StringSliceVarP(&_ethereum_tags, "tag", "t", []string{}, "Tags for the password")
	ethereumCmd.Flags().StringVarP(&_ethereum_encrypt_password, "encrypt-password", "e", "", "Password for the private key")
}

// Password

var _password_size int
var _password_tags []string
var _password_encrypt_password string

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a password",
	Run: func(cmd *cobra.Command, args []string) {
		gen.GeneratePassword(_password_size, _password_encrypt_password, _password_tags...)
	},
}

func initPasswordCmd() {
	genCmd.AddCommand(passwordCmd)
	passwordCmd.Flags().IntVarP(&_password_size, "size", "s", 64, "Size of the password")
	passwordCmd.Flags().StringSliceVarP(&_password_tags, "tag", "t", []string{}, "Tags for the password")
	passwordCmd.Flags().StringVarP(&_password_encrypt_password, "encrypt-password", "e", "", "Password for the encryption")
}

// Proto

var _proto_type string
var _proto_package string
var _proto_output_path []string

var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "Generate a proto file from an ABI read from stdin",
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateProto(os.Stdin, _proto_type, _proto_package, _proto_output_path...)
	},
}

func initProtoCmd() {
	genCmd.AddCommand(protoCmd)
	protoCmd.Flags().StringVarP(&_proto_type, "type", "t", "Contract", "Type name for the proto file")
	protoCmd.Flags().StringVarP(&_proto_package, "pkg", "p", "contract.v1", "Package name for the proto file")
	protoCmd.Flags().StringArrayVarP(&_proto_output_path, "out", "o", []string{}, "Path to the output file")
}
