package collections

import (
	"gopkg.in/yaml.v3"
)

func ParseYamlFile(contents []byte) (Collection, error) {
	collection := new(Collection)

	err := yaml.Unmarshal(contents, collection)
	if err != nil {
		return Collection{}, err
	}

	return *collection, nil
}
