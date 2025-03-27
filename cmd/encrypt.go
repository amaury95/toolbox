/*
Copyright © 2025 Amaury Diaz <amauryuh@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _encrypt_output string

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt functionalities",
	Long:  `Encrypt a text using a password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encrypt called")
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.PersistentFlags().StringVarP(&_encrypt_output, "output", "o", "encrypted", "Output file")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
