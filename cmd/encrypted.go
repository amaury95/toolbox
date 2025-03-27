/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/amaury95/toolbox/src/util"
	"github.com/spf13/cobra"
)

var _encrypted_clean_output bool

// encryptedCmd represents the encrypted command
var encryptedCmd = &cobra.Command{
	Use:   "encrypted <file>",
	Short: "Decrypt a file",
	Long:  `Decrypt a file using a password`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		util.ReadEncryptedZip(args[0], util.PromptPassword(), _encrypted_clean_output)
	},
}

func init() {
	readCmd.AddCommand(encryptedCmd)
	encryptedCmd.Flags().BoolVarP(&_encrypted_clean_output, "clean", "c", false, "Clean the output")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
