package responses

import (
	"errors"
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

func (r *DefaultResponseFormatter) Format() (string, error) {
	body, err := r.getBody()
	if err != nil {
		return "", err
	}

	return body, nil
}
