package core

import (
	"net/http"

	"github.com/pheever/stresster/sampler"
)

//Thread executing a plan
type Thread struct {
	Var        map[string]string `yaml:"var,omitempty"`
	Samplers   []sampler.Sampler `yaml:"samplers,omitempty"`
	HTTPClient *http.Client
}
