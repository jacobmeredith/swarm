package runner

import (
	"errors"
	"net/http"

	"github.com/jacobmeredith/swarm/internal/requests"
	"github.com/jacobmeredith/swarm/internal/responses"
)

type Runner struct {
	client             *http.Client
	response_formatter responses.ResponseFormatter
}

func NewRunner(client *http.Client) *Runner {
	return &Runner{
		client: client,
	}
}

func (r *Runner) SetResponseFormatter(formatter responses.ResponseFormatter) {
	r.response_formatter = formatter
}

func (r *Runner) Run(request *requests.Request) (string, error) {
	req, err := request.Build()
	if err != nil {
		return "", err
	}

	res, err := r.client.Do(req)
	if err != nil {
		return "", err
	}

	if r.response_formatter == nil {
		return "", errors.New("No response formatter provided")
	}

	r.response_formatter.SetResponse(res)
	r.response_formatter.SetRequest(req, request)

	formatted_response, err := r.response_formatter.Format()
	if err != nil {
		return "", err
	}

	return formatted_response, nil
}
