package requests

type Request struct {
	Url    string `yaml:"url" json:"url"`
	Method string `yaml:"method" json:"method"`
}
