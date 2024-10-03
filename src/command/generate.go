package command

import (
	"github.com/amaury95/toolbox/src/gen"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a password, an ethereum address, etc.",
}

func init() {
	rootCmd.AddCommand(genCmd)
	initBitcoinCmd()
	initEthereumCmd()
	initPasswordCmd()
}

// Bitcoin

var _bitcoin_tags []string
var _bitcoin_encrypt_password string

var bitcoinCmd = &cobra.Command{
	Use:   "bitcoin",
	Short: "Generate a Bitcoin private key and address",
	Run: func(cmd *cobra.Command, args []string) {
		gen.GenerateBitcoinKey(_bitcoin_encrypt_password, _bitcoin_tags...)
	},
}

func initBitcoinCmd() {
	genCmd.AddCommand(bitcoinCmd)
	bitcoinCmd.Flags().StringSliceVarP(&_bitcoin_tags, "tag", "t", []string{}, "Tags for the password")
	bitcoinCmd.Flags().StringVarP(&_bitcoin_encrypt_password, "encrypt-password", "e", "", "Password for the encryption")
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
