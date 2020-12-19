package httpsampler

import (
	"net/http"

	"github.com/pheever/stresster/sampler"
)

//HTTPSampler is a sampler
type HTTPSampler struct {
	sampler.Sampler
	HTTP struct {
		URL        string            `yaml:"url,omitempty"`
		Body       string            `yaml:"body,omitempty"`
		Method     string            `yaml:"method,omitempty"`
		Form       map[string]string `yaml:"form,omitempty"`
		Headers    map[string]string `yaml:"headers,omitempty"`
		Parameters map[string]string `yaml:"parameters,omitempty"`
	} `yaml:"http,omitempty"`
}

//Config of the sampler
type Config struct {
	AllowCookies bool `yaml:"allowCookies,omitempty"`
}

//NewHTTPClient creates a new client
func NewHTTPClient() *http.Client {
	return http.DefaultClient
}
