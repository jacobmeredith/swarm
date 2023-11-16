package responses

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Joker/hpp"
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

func (r *ResponseBuilder) buildJsonResponse(body []byte) string {
	return fmt.Sprintf("```json\n%v\n```", string(body))
}

func (r *ResponseBuilder) buildHtmlResponse(body []byte) string {
	pretty := hpp.Print(bytes.NewReader(body))
	return fmt.Sprintf("```html\n%v\n```", string(pretty))
}

func (r *ResponseBuilder) Render() (string, error) {
	body, err := r.getBody()
	if err != nil {
		return "", err
	}

	md := fmt.Sprintf("# %v [%v](%v)\n", r.req.Method, r.req.URL, r.req.URL)

	ct := r.res.Header.Get("Content-type")

	if strings.Contains(ct, "application/json") {
		json := r.buildJsonResponse(body)
		md += json
		return renderer.Render(md)
	}

	if strings.Contains(ct, "text/html") {
		md += r.buildHtmlResponse(body)
		return renderer.Render(md)
	}

	md += fmt.Sprintf("```\n%v\n```", string(body))
	return renderer.Render(md)
}
