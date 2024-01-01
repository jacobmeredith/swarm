package requests

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"slices"
	"strings"
)

type RequestCreatorOptions struct {
	Url         string
	Method      string
	ContentType string
	Body        string
	Headers     string
	Cookies     string
}

type Request struct {
	Url         string
	Method      string
	ContentType string
	Body        io.Reader
	Headers     map[string]string
	Cookies     map[string]string
}

func NewRequestConfig(options RequestCreatorOptions) (*Request, []error) {
	errors := make([]error, 0)

	url, err := options.getUrl()
	if err != nil {
		errors = append(errors, err)
	}

	method, err := options.getMethod()
	if err != nil {
		errors = append(errors, err)
	}

	content_type, err := options.getValidContentType()
	if err != nil {
		errors = append(errors, err)
	}

	request_config := &Request{
		Url:         url,
		Method:      method,
		ContentType: content_type,
		Body:        options.getBody(),
	}

	return request_config, errors
}

func (rco *RequestCreatorOptions) getUrl() (string, error) {
	if rco.Url == "" {
		return "", errors.New("No URL provided")
	}

	return rco.Url, nil
}

func (rco *RequestCreatorOptions) getMethod() (string, error) {
	if rco.Method == "" {
		return "", errors.New("No method provided")
	}

	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions, http.MethodHead}

	if !slices.Contains(methods, rco.Method) {
		return "", errors.New("Invalid method provided")
	}

	return rco.Method, nil
}

func (rco *RequestCreatorOptions) getValidContentType() (string, error) {
	if rco.ContentType == "" {
		return "", errors.New("No content type provided")
	}

	if len(strings.Split(rco.ContentType, "/")) < 2 {
		return "", errors.New("Invalid content type provided")
	}

	return rco.ContentType, nil
}

func (rco *RequestCreatorOptions) getBody() io.Reader {
	if rco.Body == "" {
		return nil
	}

	body := []byte(rco.Body)
	return bytes.NewReader(body)
}

type RequestCreator struct {
	config      *Request
	http_client *http.Client
}

func NewRequestCreator(options RequestCreatorOptions, http_client *http.Client) (*RequestCreator, error) {
	request_config, errs := NewRequestConfig(options)

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return &RequestCreator{
		config:      request_config,
		http_client: http_client,
	}, nil
}

func (rc *RequestCreator) Create() (*http.Response, error) {
	_, err := http.NewRequest(rc.config.Method, rc.config.Url, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
