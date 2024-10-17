package command

import (
	"github.com/amaury95/toolbox/src/encrypt"
	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a text",
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	initTextEncryptCmd()
}

var _text_encrypt_content string
var _text_encrypt_name string
var _text_encrypt_password string

var textEncryptCmd = &cobra.Command{
	Use:   "text",
	Short: "Encrypt a text",
	Run: func(cmd *cobra.Command, args []string) {
		encrypt.EncryptText(_text_encrypt_name, _text_encrypt_content, _text_encrypt_password)
	},
}

func initTextEncryptCmd() {
	encryptCmd.AddCommand(textEncryptCmd)
	textEncryptCmd.Flags().StringVarP(&_text_encrypt_name, "name", "n", "", "Name for the encrypted file")
	textEncryptCmd.Flags().StringVarP(&_text_encrypt_password, "password", "p", "", "Password for the encryption")
	textEncryptCmd.Flags().StringVarP(&_text_encrypt_content, "content", "c", "", "Text to encrypt")
}
