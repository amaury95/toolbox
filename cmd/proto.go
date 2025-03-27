/*
Copyright © 2025 Amaury Diaz <amauryuh@gmail.com>
*/
package cmd

import (
	"os"

	gen "github.com/amaury95/toolbox/src/generate"
	"github.com/spf13/cobra"
)

var _proto_type string
var _proto_package string

// protoCmd represents the proto command
var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "Generate a proto file",
	Long:  `Generate a proto file`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateProto(os.Stdin, _proto_type, _proto_package, _generate_output)
	},
}

func init() {
	generateCmd.AddCommand(protoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// protoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// protoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
