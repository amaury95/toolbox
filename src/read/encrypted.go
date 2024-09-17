package read

import (
	"log"

	"github.com/amaury95/toolbox/src/util"
)

func ReadEncryptedFile(file string, password string) {
	if err := util.ReadEncryptedZip(file, password); err != nil {
		log.Println("error reading encrypted zip", err)
	}
}
