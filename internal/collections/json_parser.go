package collections

import (
	"encoding/json"
)

func ParseJsonFile(contents []byte) (Collection, error) {
	collection := new(Collection)

	err := json.Unmarshal(contents, collection)
	if err != nil {
		return Collection{}, err
	}

	return *collection, nil
}
