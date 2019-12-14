package core

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//Plan is the shell element
type Plan struct {
	Name       string                 `json:"name,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Iterations uint                   `json:"iterations,omitempty"`
	Users      uint                   `json:"users,omitempty"`
	RampUp     uint                   `json:"rampup,omitempty"`
	HoldLoad   uint                   `json:"holdload,omitempty"`
	Tasks      []interface{}          `json:"tasks"`
}

//NewPlan makes a new plan
func NewPlan(name string) Plan {
	var plan Plan
	plan.Name = name
	return plan
}

//LoadPlanFromFile reads a Plan from a file
func LoadPlanFromFile(filePath string) (*Plan, error) {
	var plan Plan

	ext := filepath.Ext(filePath)
	unmarshal, supported := unmarshallers[ext]
	if !supported {
		return nil, fmt.Errorf("Unsupported plan file type: '%s'", ext)
	}

	raw, readerr := ioutil.ReadFile(filePath)
	if readerr != nil {
		return nil, readerr
	}

	err := unmarshal(raw, &plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}
