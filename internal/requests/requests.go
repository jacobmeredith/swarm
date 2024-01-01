package requests

import (
	"errors"
	"io"
	"net/http"
)

type Request struct {
	Url         string
	Method      string
	ContentType string
	Body        io.Reader
	Headers     map[string]string
	Cookies     map[string]string
}

func NewRequest(config RequestConfig) (*Request, error) {
	errs := make([]error, 0)

	url, err := config.getUrl()
	if err != nil {
		errs = append(errs, err)
	}

	method, err := config.getMethod()
	if err != nil {
		errs = append(errs, err)
	}

	content_type, err := config.getValidContentType()
	if err != nil {
		errs = append(errs, err)
	}

	return &Request{
		Url:         url,
		Method:      method,
		ContentType: content_type,
		Body:        config.getBody(),
		Headers:     config.getHeaders(),
		Cookies:     config.getCookies(),
	}, errors.Join(errs...)
}

func (r *Request) Build() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
