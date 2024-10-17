package read

import (
	"log"

	"github.com/amaury95/toolbox/src/util"
)

func ReadEncryptedFile(file string, password string, cleanOutput bool) {
	if err := util.ReadEncryptedZip(file, password, cleanOutput); err != nil {
		log.Println("error reading encrypted zip", err)
	}
}
