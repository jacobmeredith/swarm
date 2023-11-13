package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHeaders(t *testing.T) {
	headerString := "authorization:Bearer token,x-test-header:test"
	parsed := ParseHeaders(headerString)

	assert.Len(t, parsed, 2, "Length should be 2")
	assert.Equal(t, "Bearer token", parsed["authorization"])
	assert.Equal(t, "test", parsed["x-test-header"])

	headerString = ""
	parsed = ParseHeaders(headerString)

	assert.Len(t, parsed, 0)
}
