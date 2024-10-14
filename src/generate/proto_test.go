package gen_test

import (
	"os"
	"testing"

	"github.com/amaury95/toolbox/src/generate"
	"github.com/stretchr/testify/assert"
)

func TestGenerateProto(t *testing.T) {
	abi, err := os.Open("abi.json")
	assert.NoError(t, err)

	gen.GenerateProto(abi, "Test", "test", "test.proto")
}
