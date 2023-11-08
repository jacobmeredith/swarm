package collections

import (
	"testing"
)

func TestParser(t *testing.T) {
	cb := NewCollectionBuilder("test.json", func(path string) ([]byte, error) {
		return []byte(`{
			"requests": {
				"GetTest": {
					"method": "GET",
					"url": "https://google.com"
				}
			}
		}`), nil
	})

	collection, err := cb.Build()
	if err != nil {
		t.Fatalf("Could not build collection: %v", err)
	}

	testRequest := collection.Requests["GetTest"]

	if testRequest.Url != "https://google.com" {
		t.Fatal("Url is incorrect")
	}

	if testRequest.Method != "GET" {
		t.Fatal("Method is incorrect")
	}
}
