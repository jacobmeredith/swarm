package requests

type Request struct {
	Url         string `yaml:"url" json:"url"`
	Method      string `yaml:"method" json:"method"`
	ContentType string `yaml:"content-type" json:"contentType"`
	Body        string `yaml:"body" json:"body"`
}
