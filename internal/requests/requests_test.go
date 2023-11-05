package requests

import (
	"testing"
)

func TestParseHeaders(t *testing.T) {
	headerString := "authorization:Bearer token,x-test-header:test"
	parsed := ParseHeaders(headerString)

	if len(parsed) != 2 {
		t.Fatalf("Length should be 2, instead recieved: %v", len(parsed))
	}
}
