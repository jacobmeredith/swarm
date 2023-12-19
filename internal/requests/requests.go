package requests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jacobmeredith/swarm/internal/responses"
)

type Request struct {
	Url         string            `yaml:"url" json:"url"`
	Method      string            `yaml:"method" json:"method"`
	ContentType string            `yaml:"content-type" json:"contentType"`
	Body        string            `yaml:"body" json:"body"`
	Headers     map[string]string `yaml:"headers" json:"headers"`
	Cookies     map[string]string `yaml:"cookies" json:"cookies"`
}

func ParseKeyValuePairString(text string, pair_delimeter string, value_delimeter string) map[string]string {
	values_map := make(map[string]string)

	parts := strings.Split(text, pair_delimeter)

	if len(parts) == 1 && parts[0] == "" {
		return values_map
	}

	for _, p := range parts {
		pSplit := strings.Split(p, value_delimeter)
		values_map[pSplit[0]] = pSplit[1]
	}

	return values_map
}

func NewRequest(method, url, contentType, body string, headers map[string]string, cookies map[string]string) *Request {
	return &Request{
		Method:      method,
		Url:         url,
		ContentType: contentType,
		Body:        body,
		Headers:     headers,
		Cookies:     cookies,
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

func (r *Request) buildCookies(req *http.Request) {
	for key, value := range r.Cookies {
		req.AddCookie(&http.Cookie{Name: key, Value: value})
	}
}

func (r *Request) Run() error {
	body := r.buildBody()
	req, _ := http.NewRequest(r.Method, r.Url, body)
	r.buildHeaders(req)
	r.buildCookies(req)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	resBuilder := responses.NewResponseBuilder(req, res)
	out, err := resBuilder.Render()
	if err != nil {
		return err
	}

	fmt.Print(out)

	return nil
}
