package responses

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jacobmeredith/swarm/internal/requests"
)

type ResponseFormatter interface {
	SetResponse(response *http.Response)
	SetRequest(request *http.Request, custom *requests.Request)
	Format() (string, error)
}

type DefaultResponseFormatter struct {
	native_response *http.Response
	native_request  *http.Request
	request         *requests.Request
}

func NewDefaultResponseFormatter() ResponseFormatter {
	return &DefaultResponseFormatter{}
}

func (r *DefaultResponseFormatter) SetResponse(response *http.Response) {
	r.native_response = response
}

func (r *DefaultResponseFormatter) SetRequest(request *http.Request, custom *requests.Request) {
	r.native_request = request
	r.request = custom
}

func (r *DefaultResponseFormatter) getBody() (string, error) {
	if r.native_response == nil {
		return "", errors.New("No response provided")
	}

	body, err := io.ReadAll(r.native_response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (r *DefaultResponseFormatter) getHeaders() string {
	if r.native_response == nil {
		return ""
	}

	headers := ""
	for key, value := range r.native_response.Header {
		headers += key + ": " + value[0] + "\n"
	}

	return headers
}

func (r *DefaultResponseFormatter) Format() (string, error) {
	response := ""

	response += fmt.Sprintf("Request:\n%s %s\n", r.native_request.Method, r.native_request.URL.String())
	response += fmt.Sprintf("Status: %s\n\n", r.native_response.Status)

	headers := r.getHeaders()
	response += fmt.Sprintf("Response:\nHeaders:\n%s\n", headers)

	body, err := r.getBody()
	if err != nil {
		return "", err
	}

	response += fmt.Sprintf("body:\n%s\n", body)

	return response, nil
}
