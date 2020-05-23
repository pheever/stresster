package core

import "time"

//SampleResult describes the result
type SampleResult struct {
	Started time.Time
	Lasted  time.Duration
	ID      string
}

//ResultStream is a channel of SampleResults
type ResultStream chan<- SampleResult
