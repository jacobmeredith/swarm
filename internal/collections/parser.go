package collections

import (
	"errors"
	"strings"
)

type FileReader = func(string) ([]byte, error)

type CollectionBuilder struct {
	path     string
	fileType string
	reader   FileReader
}

func NewCollectionBuilder(path string, reader FileReader) *CollectionBuilder {
	typeSplits := strings.Split(path, ".")
	fileType := typeSplits[len(typeSplits)-1]

	return &CollectionBuilder{
		path:     path,
		fileType: fileType,
		reader:   reader,
	}
}

func (p *CollectionBuilder) getFileContents() ([]byte, error) {
	b, err := p.reader(p.path)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b *CollectionBuilder) Build() (*Collection, error) {
	var collection Collection
	var err error

	contents, err := b.getFileContents()
	if err != nil {
		return &collection, err
	}

	switch b.fileType {
	case "yaml":
		collection, err = ParseYamlFile(contents)
		return &collection, err
	case "json":
		collection, err = ParseJsonFile(contents)
		return &collection, err
	default:
		return &collection, errors.New("File type not supported")
	}
}
