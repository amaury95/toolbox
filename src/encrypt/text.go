package encrypt

import (
	"io"

	"github.com/amaury95/toolbox/src/util"
)

func EncryptText(name, password string, reader io.Reader) error {
	content, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	return util.CreateEncryptedZip(name, string(content), password)
}
