package requests

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func getFileType(filename string) string {
	parts := strings.Split(filename, ".")
	return parts[len(parts)-1]
}

type CollectionRequest struct {
	Url         string            `yaml:"url" json:"url"`
	Method      string            `yaml:"method" json:"method"`
	ContentType string            `yaml:"content_type" json:"content_type"`
	Body        string            `yaml:"body" json:"body"`
	Headers     map[string]string `yaml:"headers" json:"headers"`
	Cookies     map[string]string `yaml:"cookies" json:"cookies"`
}

type Collection struct {
	Requests map[string]CollectionRequest `yaml:"requests" json:"requests"`
}

func NewCollection(directory, filename, name string) (*Collection, error) {
	path := fmt.Sprintf("%s/%s", directory, filename)

	file_contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	file_type := getFileType(filename)
	collection := new(Collection)

	switch file_type {
	case "yaml":
		err = yaml.Unmarshal(file_contents, &collection)
		if err != nil {
			return nil, err
		}
	case "json":
		err = json.Unmarshal(file_contents, &collection)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unsupported file type: %s", file_type)
	}

	return collection, nil
}
