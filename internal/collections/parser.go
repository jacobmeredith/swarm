package collections

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type CollectionBuilder struct {
	path     string
	fileType string
}

func NewCollectionBuilder(path string) *CollectionBuilder {
	fmt.Printf("%+v\n", path)
	typeSplits := strings.Split(path, ".")
	fileType := typeSplits[len(typeSplits)-1]

	return &CollectionBuilder{
		path:     path,
		fileType: fileType,
	}
}

func (p *CollectionBuilder) getFileContents() ([]byte, error) {
	b, err := os.ReadFile(p.path)
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
		fmt.Printf("%+v\n", collection)
		return &collection, err
	case "json":
		collection, err = ParseJsonFile(contents)
		fmt.Printf("%+v\n", collection)
		return &collection, err
	default:
		return &collection, errors.New("File type not supported")
	}
}
