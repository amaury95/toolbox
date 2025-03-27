/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
