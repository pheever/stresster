package sampler

//Sampler executes and generates a sample
type Sampler struct {
	Name       string                 `json:"name,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Config     map[string]interface{} `json:"config,omitempty"`
	Iterations int                    `json:"iterations,omitempty"`
}
