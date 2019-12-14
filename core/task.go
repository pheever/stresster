package core

//Task core struct
type coreTask struct {
	Name     string `json:"name"`
	When     string `json:"when"`
	Emit     string `json:"emit"`
	Listen   string `json:"listen"`
	Register string `json:"register"`
	Assert   string `json:"assert"`
}

//Task to execute
type Task struct {
	coreTask
}

//GetName gets the Tasks name
func (task *Task) GetName() string {
	return task.Name
}

//GetWhen gets the Tasks when condition
func (task *Task) GetWhen() string {
	return task.When
}

//GetEmit gets the Tasks emit
func (task *Task) GetEmit() string {
	return task.Emit
}

//GetListen gets the Tasks listen event
func (task *Task) GetListen() string {
	return task.Listen
}

//GetRegister gets the Tasks register value
func (task *Task) GetRegister() string {
	return task.Register
}

//GetAssert gets the Tasks assert condition
func (task *Task) GetAssert() string {
	return task.Assert
}
