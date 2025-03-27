/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	gen "github.com/amaury95/toolbox/src/generate"
	"github.com/amaury95/toolbox/src/util"
	"github.com/spf13/cobra"
)

var _password_length int
var _password_include_uppercase bool
var _password_include_numbers bool
var _password_include_symbols bool

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		gen.GeneratePassword(_password_length, util.GeneratePasswordOptions{
			IncludeUppercase: _password_include_uppercase,
			IncludeNumbers:   _password_include_numbers,
			IncludeSymbols:   _password_include_symbols,
		}, _generate_output, util.PromptPasswordConfirm(), _generate_tags...)
	},
}

func init() {
	generateCmd.AddCommand(passwordCmd)
	passwordCmd.Flags().IntVarP(&_password_length, "length", "l", 16, "Length of the password")
	passwordCmd.Flags().BoolVarP(&_password_include_uppercase, "uppercase", "u", true, "Include uppercase letters in the password")
	passwordCmd.Flags().BoolVarP(&_password_include_numbers, "numbers", "n", true, "Include numbers in the password")
	passwordCmd.Flags().BoolVarP(&_password_include_symbols, "symbols", "s", true, "Include symbols in the password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
