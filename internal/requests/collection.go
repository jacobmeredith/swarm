package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func getFileType(filename string) string {
	parts := strings.Split(filename, ".")
	return parts[len(parts)-1]
}

type ParserFunction = func(in []byte, out interface{}) error

func parser(contents []byte, parser ParserFunction) (map[string]CollectionRequest, error) {
	collection := new(CollectionFile)

	err := parser(contents, &collection)
	if err != nil {
		return nil, err
	}

	return collection.Requests, nil
}

type CollectionFile struct {
	Requests map[string]CollectionRequest `yaml:"requests" json:"requests"`
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
	FileType  string
	Requests  map[string]CollectionRequest `yaml:"requests" json:"requests"`
}

func NewCollection(directory, filename string) (*Collection, error) {
	collection := &Collection{
		Directory: directory,
		Filename:  filename,
		Path:      fmt.Sprintf("%s/%s", directory, filename),
		FileType:  getFileType(filename),
	}

	contents, err := collection.readFile()
	if err != nil {
		return nil, err
	}

	requests, err := collection.parseFile(contents)
	if err != nil {
		return nil, err
	}

	collection.Requests = requests

	return collection, nil
}

func (c *Collection) readFile() ([]byte, error) {
	contents, err := os.ReadFile(c.Path)
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (c *Collection) parseFile(contents []byte) (map[string]CollectionRequest, error) {
	switch c.FileType {
	case "yaml":
		return parser(contents, yaml.Unmarshal)
	case "json":
		return parser(contents, json.Unmarshal)
	}

	return nil, errors.New("Invalid file type provided")
}

func (c *Collection) GetCollectionRequest(name string) (*CollectionRequest, error) {
	request, ok := c.Requests[name]
	if !ok {
		return nil, errors.New("Request not found")
	}

	return &request, nil
}

func (c *Collection) TransformRequest(name string) (*Request, error) {
	request, err := c.GetCollectionRequest(name)
	if err != nil {
		return nil, err
	}

	return &Request{
		Url:         request.Url,
		Method:      request.Method,
		ContentType: request.ContentType,
		Body:        strings.NewReader(request.Body),
		Headers:     request.Headers,
		Cookies:     request.Cookies,
	}, nil
}
