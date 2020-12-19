package sampler

//Sampler executes and generates a sample
type Sampler struct {
	Tags       []string               `yaml:"tags,omitempty"`
	Name       string                 `yaml:"name,omitempty"`
	Type       string                 `yaml:"type,omitempty"`
	Config     map[string]interface{} `yaml:"config,omitempty"`
	Iterations int                    `yaml:"iterations,omitempty"`
}
