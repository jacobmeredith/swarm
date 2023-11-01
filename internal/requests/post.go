package requests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Post(url string, ct string, body string) error {
	b := []byte(body)
	bReader := bytes.NewReader(b)

	resp, err := http.Post(url, ct, bReader)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	rBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(rBody))

	return nil
}
