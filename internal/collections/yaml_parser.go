package collections

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YamlParser struct {
	path string
}

func NewYamlParser(path string) *YamlParser {
	return &YamlParser{
		path: path,
	}
}

func (p *YamlParser) getFileContents() ([]byte, error) {
	b, err := os.ReadFile(p.path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (p *YamlParser) ParseFile() (Collection, error) {
	contents, err := p.getFileContents()
	if err != nil {
		return Collection{}, err
	}

	collection := new(Collection)

	err = yaml.Unmarshal(contents, collection)
	if err != nil {
		return Collection{}, err
	}

	return *collection, nil
}
