package gen

import (
	"log"
	"strings"

	"github.com/amaury95/toolbox/src/util"
)

func GeneratePassword(size int, opts util.GeneratePasswordOptions, output string, encryptPassword string, tags ...string) {
	password := util.GenerateRandomPassword(size, opts)
	log.Println("generated password:", password)

	if encryptPassword != "" {
		if output == "" {
			output = strings.Join(append(tags, "password"), "_")
		}
		if err := util.CreateEncryptedZip(output, password, encryptPassword); err != nil {
			log.Fatal(err)
		}
	}
}
