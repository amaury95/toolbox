/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	gen "github.com/amaury95/toolbox/src/generate"
	"github.com/amaury95/toolbox/src/util"
	"github.com/spf13/cobra"
)

var _address_protocol string

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateEthereumKey(_generate_output, util.PromptPasswordConfirm(), _address_protocol, _generate_tags...)
	},
}

func init() {
	generateCmd.AddCommand(addressCmd)
	addressCmd.Flags().StringVarP(&_address_protocol, "protocol", "p", "ethereum", "Protocol for the address")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
