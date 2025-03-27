/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	gen "github.com/amaury95/toolbox/src/generate"
	"github.com/amaury95/toolbox/src/util"
	"github.com/spf13/cobra"
)

// mnemonicCmd represents the mnemonic command
var mnemonicCmd = &cobra.Command{
	Use:   "mnemonic",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateMnemonic(_generate_output, util.PromptPasswordConfirm(), _generate_tags...)
	},
}

func init() {
	generateCmd.AddCommand(mnemonicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mnemonicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mnemonicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
