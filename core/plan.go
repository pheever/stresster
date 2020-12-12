package core

import "github.com/pheever/stresster/sampler"

//Plan to execute
type Plan struct {
	Name     string
	Samplers []sampler.Sampler
}
