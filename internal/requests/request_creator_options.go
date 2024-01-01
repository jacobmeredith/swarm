package requests

import (
	"errors"
	"io"
	"net/http"
	"slices"
	"strings"
)

func splitKeyValuePairString(text string, delimete string, separator string) map[string]string {
	kv := make(map[string]string)
	parts := strings.Split(text, delimete)

	for _, part := range parts {
		kvPair := strings.Split(part, separator)
		kv[kvPair[0]] = kvPair[1]
	}

	return kv
}

type RequestCreatorOptions struct {
	Url         string
	Method      string
	ContentType string
	Body        string
	Headers     string
	Cookies     string
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

	return strings.NewReader(rco.Body)
}

func (rco *RequestCreatorOptions) getHeaders() map[string]string {
	if rco.Headers == "" {
		return nil
	}

	return splitKeyValuePairString(rco.Headers, ",", ":")
}

func (rco *RequestCreatorOptions) getCookies() map[string]string {
	if rco.Cookies == "" {
		return nil
	}

	return splitKeyValuePairString(rco.Cookies, ",", ":")
}
