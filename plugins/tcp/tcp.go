package tcpmodule

import "github.com/pheever/stresster/core"

//TCPModule implements tcp testing module
type TCPModule struct {
	core.Task `yaml:",inline"`
	TCPModule struct {
		Protocol string `yaml:"protocol"`
		Host     string `yaml:"url"`
		Port     int    `yaml:"port"`
	} `yaml:"tcp"`
}
