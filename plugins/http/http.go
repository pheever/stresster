package httpmodule

import "github.com/pheever/stresster/core"

//HTTPModule implements http testing module
type HTTPModule struct {
	core.Task  `yaml:",inline"`
	HTTPModule struct {
		Protocol string `yaml:"protocol"`
		Host     string `yaml:"url"`
		Port     int    `yaml:"port"`
	} `yaml:"http"`
}

//Start module execution
func (h *HTTPModule) Start() {

}
