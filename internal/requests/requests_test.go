package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitKeyValuePairString(t *testing.T) {
	headerString := "authorization:Bearer token,x-test-header:test"
	parsed := splitKeyValuePairString(headerString, ",", ":")

	assert.Len(t, parsed, 2, "Length should be 2")
	assert.Equal(t, "Bearer token", parsed["authorization"])
	assert.Equal(t, "test", parsed["x-test-header"])

	headerString = ""
	parsed = splitKeyValuePairString(headerString, ",", ":")

	assert.Len(t, parsed, 0)
}
