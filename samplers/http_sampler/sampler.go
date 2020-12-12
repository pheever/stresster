package httpsampler

import (
	"github.com/pheever/stresster/sampler"
	"gopkg.in/yaml.v3"
)

//Config for http sampler
type Config struct {
	URL     string            `json:"url,omitempty"`
	Method  string            `json:"method,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

var configurations map[string]*Config = map[string]*Config{}

//Execute the sampler based on a config
func Execute(smplr sampler.Sampler) *sampler.Sample {
	if _, ok := configurations[smplr.Name]; !ok {
		c, e := processConfig(smplr.Config)
		if e != nil {
			return nil
		}
		configurations[smplr.Name] = c
	}
	return nil
}

//processConfig provided
func processConfig(raw map[string]interface{}) (*Config, error) {
	var c Config
	out, err := yaml.Marshal(raw)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(out, &c)
	return &c, err
}
