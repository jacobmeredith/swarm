package requests

type Request struct {
	Name   string `yaml:"name"`
	Url    string `yaml:"url"`
	Method string `yaml:"method"`
}
