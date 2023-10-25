package collections

import (
	"fmt"

	"github.com/jacobmeredith/swarm/internal/requests"
)

type Collection struct {
	Requests map[string]requests.Request `yaml:"requests"`
}

func RunCollectionRequest(directory, filename, name string) error {
	yp := NewYamlParser(directory + "/" + filename + ".yaml")

	collection, err := yp.ParseFile()
	if err != nil {
		return err
	}

	request := collection.Requests[name]

	switch request.Method {
	case "GET":
		requests.Get(request.Url)
	default:
		fmt.Println("Method not supported")
	}

	return nil
}
