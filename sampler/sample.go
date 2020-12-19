package sampler

import "time"

//Sample of a sampler execution
type Sample struct {
	StartTime time.Time     `yaml:"start_time,omitempty"`
	Duration  time.Duration `yaml:"duration,omitempty"`
	Response  []byte        `yaml:"response,omitempty"`
	Failed    bool          `yaml:"failed,omitempty"`
	Error     error         `yaml:"error,omitempty"`
}
