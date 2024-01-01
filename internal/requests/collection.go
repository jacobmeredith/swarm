package requests

import (
	"fmt"
	"strings"
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
	Directory string
	Filename  string
	Path      string
	Requests  map[string]CollectionRequest `yaml:"requests" json:"requests"`
}

func NewCollection(directory, filename, name string) *Collection {
	return &Collection{
		Directory: directory,
		Filename:  filename,
		Path:      fmt.Sprintf("%s/%s", directory, filename),
	}
}
