package collections

import (
	"errors"
	"fmt"

	"github.com/jacobmeredith/swarm/internal/requests"
)

type Collection struct {
	Requests map[string]requests.Request `yaml:"requests" json:"requests"`
}

func RunCollectionRequest(directory, filename, name string) error {
	cb := NewCollectionBuilder(fmt.Sprintf("%v/%v", directory, filename))

	collection, err := cb.Build()
	if err != nil {
		return err
	}

	request := collection.Requests[name]

	switch request.Method {
	case "GET":
		return requests.Get(request.Url)
	case "POST":
		return requests.Post(request.Url, request.ContentType, request.Body)

	default:
		return errors.New("Method not supported")
	}
}
