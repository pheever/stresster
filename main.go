package main

import (
	"fmt"
	"io/ioutil"

	"github.com/pheever/stresster/core"
	"gopkg.in/yaml.v3"
)

func main() {
	_, e := loadFile("examples/config.yml")
	if e != nil {
		fmt.Println("Failed to load plan")
	}
	fmt.Println("Loaded")
}

func loadFile(path string) (interface{}, error) {
	var plan core.Plan
	data, err := ioutil.ReadFile(path)
	yaml.Unmarshal(data, &plan)
	return plan, err
}
