package plan

import (
	"io/ioutil"
	"os"

	"github.com/pheever/stresster/sampler"
	httpsampler "github.com/pheever/stresster/samplers/http"
	"gopkg.in/yaml.v3"
)

//Plan to execute
type Plan struct {
	Name          string            `yaml:"name,omitempty"`
	Env           []string          `yaml:"env,omitempty"`
	Tags          []string          `yaml:"tags,omitempty"`
	Var           map[string]string `yaml:"var,omitempty"`
	Global        map[string]string `yaml:"global,omitempty"`
	Configuration Configuration     `yaml:"configuration,omitempty"`
	Samplers      []sampler.Sampler `yaml:"samplers,omitempty"`
}

//Configuration of the plan
type Configuration struct {
	HTTPSampler httpsampler.Config `yaml:"httpsampler,omitempty"`
}

//Load from file
func Load(path string) (plan *Plan, err error) {
	plan, err = loadFile(path)
	if err != nil {
		return
	}
	//fetch env vars to plan
	for _, v := range plan.Env {
		if val, e := os.LookupEnv(v); e {
			plan.Global[v] = val
		}
	}
	return
}

func loadFile(path string) (plan *Plan, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &plan)
	return
}
