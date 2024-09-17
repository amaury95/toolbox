package command

import (
	"log"
	"strings"

	"github.com/amaury95/tools/src/util"
	"github.com/spf13/cobra"
)

var _password_size int
var _password_tags []string
var _password_encrypt_password string

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a password",
	Run: func(cmd *cobra.Command, args []string) {
		generatePassword(_password_size, _password_encrypt_password, _password_tags...)
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.Flags().IntVarP(&_password_size, "size", "s", 64, "Size of the password")
	passwordCmd.Flags().StringSliceVarP(&_password_tags, "tag", "t", []string{}, "Tags for the password")
	passwordCmd.Flags().StringVarP(&_password_encrypt_password, "encrypt-password", "e", "", "Password for the encryption")
}

func generatePassword(size int, encryptPassword string, tags ...string) {
	password, err := util.GenerateRandomPassword(size)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("generated password:", password)

	if encryptPassword != "" {
		fileName := strings.Join(append(tags, "password"), "_")
		if err := util.CreateEncryptedZip(fileName, password, encryptPassword); err != nil {
			log.Fatal(err)
		}
	}
}
