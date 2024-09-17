package gen

import (
	"log"
	"strings"

	"github.com/amaury95/toolbox/src/util"
)

func GeneratePassword(size int, encryptPassword string, tags ...string) {
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
