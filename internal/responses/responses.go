package responses

import (
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/glamour"
)

var renderer, _ = glamour.NewTermRenderer(
	glamour.WithAutoStyle(),
	glamour.WithWordWrap(100),
	glamour.WithBaseURL("white"),
	glamour.WithStylesFromJSONBytes([]byte(`{
			"link": {
				"color": "white",
				"underline": true,
				"block_prefix": "(",
				"block_suffix": ")"
			},
			"link_text": {
				"color": "white",
				"bold": true
			}
		}`)),
)

type ResponseBuilder struct {
	req *http.Request
	res *http.Response
}

func NewResponseBuilder(req *http.Request, res *http.Response) *ResponseBuilder {
	return &ResponseBuilder{
		req: req,
		res: res,
	}
}

func (r *ResponseBuilder) getBody() ([]byte, error) {
	body, err := io.ReadAll(r.res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (r *ResponseBuilder) Render() (string, error) {
	body, err := r.getBody()
	if err != nil {
		return "", err
	}

	md := fmt.Sprintf("# %v [%v](%v)\n", r.req.Method, r.req.URL, r.req.URL)
	md += fmt.Sprintf("```\n%v\n```", string(body))

	return renderer.Render(md)
}
