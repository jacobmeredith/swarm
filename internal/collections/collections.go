package collections

import (
	"errors"
	"fmt"
	"os"

	"github.com/jacobmeredith/swarm/internal/requests"
)

type Collection struct {
	Requests map[string]requests.Request `yaml:"requests" json:"requests"`
}

func RunCollectionRequest(directory, filename, name string) error {
	path := fmt.Sprintf("%v/%v", directory, filename)
	contents, err := os.ReadFile(path)
	if err != nil {
		return errors.New("Unable to read file")
	}

	cb := NewCollectionBuilder(filename, contents)

	collection, err := cb.Build()
	if err != nil {
		return err
	}

	request := collection.Requests[name]

	if request.Url == "" {
		return errors.New("Request name is invalid")
	}

	return request.Run()
}
