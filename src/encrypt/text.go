package encrypt

import (
	"github.com/amaury95/toolbox/src/util"
)

func EncryptText(name, password string, content string) error {
	return util.CreateEncryptedZip(name, content, password)
}
