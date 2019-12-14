package core

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//corePlan is the core Plan element
type corePlan struct {
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
	Iterations uint                   `json:"iterations"`
	Users      uint                   `json:"users"`
	RampUp     uint                   `json:"rampup"`
	HoldLoad   uint                   `json:"holdload"`
	Tasks      []interface{}          `json:"tasks"`
}

//Plan is the shell element
type Plan struct {
	corePlan
}

//NewPlan makes a new plan
func NewPlan(name string) Plan {
	var plan Plan
	plan.Name = name
	return plan
}

//LoadPlanFromFile reads a Plan from a file
func LoadPlanFromFile(filePath string) (*Plan, error) {
	var cp corePlan

	ext := filepath.Ext(filePath)
	unmarshal, supported := unmarshallers[ext]
	if !supported {
		return nil, fmt.Errorf("Unsupported plan file type: '%s'", ext)
	}

	raw, readerr := ioutil.ReadFile(filePath)
	if readerr != nil {
		return nil, readerr
	}

	err := unmarshal(raw, &cp)
	if err != nil {
		return nil, err
	}

	return &Plan{corePlan: cp}, nil
}

//GetName gets the Plans name
func (plan *Plan) GetName() string {
	return plan.Name
}

//GetProperties gets the plans properties
func (plan *Plan) GetProperties() map[string]interface{} {
	return plan.Properties
}

//SetProperty sets the value of a property
func (plan *Plan) SetProperty(key string, value interface{}) {
	plan.Properties[key] = value
}

//GetIterations gets the Plans iterations
func (plan *Plan) GetIterations() int {
	return int(plan.Iterations)
}

//SetIterations sets the Plans iterations
func (plan *Plan) SetIterations(iterations int) {
	plan.Iterations = uint(iterations)
}

//GetUsers gets the Plans users
func (plan *Plan) GetUsers() int {
	return int(plan.Users)
}

//SetUsers sets the Plans users
func (plan *Plan) SetUsers(users int) {
	plan.Users = uint(users)
}

//GetRampUp gets the Plans ramp up
func (plan *Plan) GetRampUp() int {
	return int(plan.RampUp)
}

//SetRampUp sets the Plans ramp up
func (plan *Plan) SetRampUp(rampUp int) {
	plan.RampUp = uint(rampUp)
}

//GetHoldLoad gets the plans hold load
func (plan *Plan) GetHoldLoad() int {
	return int(plan.HoldLoad)
}

//SetHoldLoad sets the plans hold load
func (plan *Plan) SetHoldLoad(holdLoad int) {
	plan.HoldLoad = uint(holdLoad)
}
