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

func NewRequest(options RequestCreatorOptions) (*Request, error) {
	errs := make([]error, 0)

	url, err := options.getUrl()
	if err != nil {
		errs = append(errs, err)
	}

	method, err := options.getMethod()
	if err != nil {
		errs = append(errs, err)
	}

	content_type, err := options.getValidContentType()
	if err != nil {
		errs = append(errs, err)
	}

	return &Request{
		Url:         url,
		Method:      method,
		ContentType: content_type,
		Body:        options.getBody(),
		Headers:     options.getHeaders(),
		Cookies:     options.getCookies(),
	}, errors.Join(errs...)
}

func (r *Request) Build() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
