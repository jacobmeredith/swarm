package collections

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func ParseYamlFile(contents []byte) (Collection, error) {
	collection := new(Collection)

	err := yaml.Unmarshal(contents, collection)
	if err != nil {
		fmt.Println(err)
		return Collection{}, err
	}

	fmt.Println(collection)

	return *collection, nil
}
