package command

import (
	"log"
	"os"

	"github.com/amaury95/toolbox/src/encrypt"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt data",
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	initTextEncryptCmd()
}

var _text_encrypt_name string
var _text_encrypt_password string

var textEncryptCmd = &cobra.Command{
	Use:   "content",
	Short: "Encrypt stdin",
	Run: func(cmd *cobra.Command, args []string) {
		if err := encrypt.EncryptText(_text_encrypt_name, _text_encrypt_password, os.Stdin); err != nil {
			log.Fatal(err)
		}
	},
}

func initTextEncryptCmd() {
	encryptCmd.AddCommand(textEncryptCmd)
	textEncryptCmd.Flags().StringVarP(&_text_encrypt_name, "name", "n", "", "Name for the encrypted file")
	textEncryptCmd.Flags().StringVarP(&_text_encrypt_password, "password", "p", "", "Password for the encryption")
}
