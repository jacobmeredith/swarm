package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestJsonParser(t *testing.T) {
	cb := NewCollectionBuilder("test.json", []byte(`{
			"requests": {
				"GetTest": {
					"method": "GET",
					"url": "https://google.com",
					"headers": {
						"Content-Type": "application/json"
					}
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
	assert.Equal(t, "application/json", testRequest.Headers["Content-Type"])
}

func TestYamlParser(t *testing.T) {
	yamlDoc, err := yaml.Marshal(map[string]interface{}{
		"requests": map[string]interface{}{
			"GetTest": map[string]interface{}{
				"method": "GET",
				"url":    "https://google.com",
				"headers": map[string]interface{}{
					"Content-Type": "application/json",
				},
			},
		},
	})

	assert.NoError(t, err)

	cb := NewCollectionBuilder("test.yaml", yamlDoc)

	collection, err := cb.Build()
	if err != nil {
		t.Fatalf("Could not build collection: %v", err)
	}

	testRequest := collection.Requests["GetTest"]

	assert.Equal(t, "https://google.com", testRequest.Url)
	assert.Equal(t, "GET", testRequest.Method)
	assert.Equal(t, "application/json", testRequest.Headers["Content-Type"])
}
