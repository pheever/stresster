package core

//Sampler interface
type Sampler interface {
	Start()
	Stop()
	Done() Done
	Emit() ResultStream
}

//Done signals the Sampler is done
type Done <-chan bool
