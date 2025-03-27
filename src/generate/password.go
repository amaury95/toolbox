package gen

import (
	"log"
	"strings"

	"github.com/amaury95/toolbox/src/util"
)

func GeneratePassword(size int, output string, encryptPassword string, tags ...string) {
	password, err := util.GenerateRandomPassword(size)
	if err != nil {
		log.Fatal(err)
	}
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
