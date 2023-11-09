package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonParser(t *testing.T) {
	cb := NewCollectionBuilder("test.json", []byte(`{
			"requests": {
				"GetTest": {
					"method": "GET",
					"url": "https://google.com"
				}
			}
		}`),
	)

	collection, err := cb.Build()
	if err != nil {
		t.Fatalf("Could not build collection: %v", err)
	}

	testRequest := collection.Requests["GetTest"]

	assert.Equal(t, "https://google.com", testRequest.Url)
	assert.Equal(t, "GET", testRequest.Method)
}
