package requests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Url         string            `yaml:"url" json:"url"`
	Method      string            `yaml:"method" json:"method"`
	ContentType string            `yaml:"content-type" json:"contentType"`
	Body        string            `yaml:"body" json:"body"`
	Headers     map[string]string `yaml:"headers" json:"headers"`
}

func NewRequest(method, url, contentType, body string) *Request {
	return &Request{
		Method:      method,
		Url:         url,
		ContentType: contentType,
		Body:        body,
	}
}

func (r *Request) buildBody() io.Reader {
	body := []byte(r.Body)
	br := bytes.NewReader(body)
	return br
}

func (r *Request) buildHeaders(req *http.Request) {
	req.Header.Set("content-type", r.ContentType)

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
}

func (r *Request) getResBody(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (r *Request) Run() error {
	body := r.buildBody()
	req, _ := http.NewRequest(r.Method, r.Url, body)
	r.buildHeaders(req)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	resBody, err := r.getResBody(res)
	if err != nil {
		return err

	}

	fmt.Println(string(resBody))

	return nil
}
