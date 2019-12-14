package debug

import (
	"fmt"

	"github.com/pheever/stresster/core"
)

type coreDebug struct {
	Message string `json:"msg"`
}

//Debug module defines a debugger, used to print values to help debugging
type Debug struct {
	core.Task
	Debug coreDebug `json:"debug"`
}

//Start the debugging module
func (module *Debug) Start() {
	fmt.Printf("Debug:\n%s\n*******\n", module.Debug.Message)
}

//SetDebugMessage sets the debugging message
func (module *Debug) SetDebugMessage(message string) {
	module.Debug.Message = message
}
