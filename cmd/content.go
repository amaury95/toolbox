/*
Copyright © 2025 Amaury Diaz <amauryuh@gmail.com>
*/
package cmd

import (
	"github.com/amaury95/toolbox/src/encrypt"
	"github.com/amaury95/toolbox/src/util"
	"github.com/spf13/cobra"
)

// contentCmd represents the content command
var contentCmd = &cobra.Command{
	Use:   "content <text>",
	Short: "Encrypt a text",
	Long:  `Encrypt a text using a password`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		encrypt.EncryptText(_encrypt_output, util.PromptPasswordConfirm(), args[0])
	},
}

func init() {
	encryptCmd.AddCommand(contentCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
