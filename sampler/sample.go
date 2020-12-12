package sampler

import "time"

//Sample of a sampler execution
type Sample struct {
	StartTime time.Time     `json:"start_time,omitempty"`
	Duration  time.Duration `json:"duration,omitempty"`
	Response  []byte        `json:"response,omitempty"`
	Failed    bool          `json:"failed,omitempty"`
	Error     error         `json:"error,omitempty"`
}
