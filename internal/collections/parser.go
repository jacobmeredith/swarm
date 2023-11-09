package collections

import (
	"errors"
	"strings"
)

type CollectionBuilder struct {
	path     string
	fileType string
	contents []byte
}

func NewCollectionBuilder(path string, contents []byte) *CollectionBuilder {
	typeSplits := strings.Split(path, ".")
	fileType := typeSplits[len(typeSplits)-1]

	return &CollectionBuilder{
		path:     path,
		fileType: fileType,
		contents: contents,
	}
}

func (b *CollectionBuilder) Build() (*Collection, error) {
	var collection Collection
	var err error

	switch b.fileType {
	case "yaml":
		collection, err = ParseYamlFile(b.contents)
		return &collection, err
	case "json":
		collection, err = ParseJsonFile(b.contents)
		return &collection, err
	default:
		return &collection, errors.New("File type not supported")
	}
}
