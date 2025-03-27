package util_test

import (
	"testing"

	"github.com/amaury95/toolbox/src/util"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomPassword(t *testing.T) {
	password := util.GenerateRandomPassword(64, util.GeneratePasswordOptions{
		IncludeUppercase: true,
		IncludeNumbers:   true,
		IncludeSymbols:   true,
	})
	assert.Equal(t, 64, len(password))
}
