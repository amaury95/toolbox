package command

import (
	"github.com/amaury95/toolbox/src/read"
	"github.com/spf13/cobra"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a password, an ethereum address, etc.",
}

func init() {
	rootCmd.AddCommand(readCmd)
	initEncryptedCmd()
}

// Encrypted

var _encrypted_file string
var _encrypted_password string

var encryptedCmd = &cobra.Command{
	Use:   "encrypted",
	Short: "Read an encrypted file",
	Run: func(cmd *cobra.Command, args []string) {
		read.ReadEncryptedFile(_encrypted_file, _encrypted_password)
	},
}

func initEncryptedCmd() {
	readCmd.AddCommand(encryptedCmd)
	encryptedCmd.Flags().StringVarP(&_encrypted_file, "file", "f", "", "File to read")
	encryptedCmd.Flags().StringVarP(&_encrypted_password, "password", "p", "", "Password to decrypt the file")
}
