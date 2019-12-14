package core

//Task to execute
type Task struct {
	Name     *string `json:"name"`
	When     *string `json:"when,omitempty"`
	Emit     *string `json:"emit,omitempty"`
	Listen   *string `json:"listen,omitempty"`
	Register *string `json:"register,omitempty"`
	Assert   *string `json:"assert,omitempty"`
}
