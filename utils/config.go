package stresster

type Task struct {
	Name     *string `json:"name"`
	When     *string `json:"when,omitempty"`
	Emit     *string `json:"emit,omitempty"`
	Listen   *string `json:"listen,omitempty"`
	Assert   *string `json:"assert,omitempty"`
	Register *string `json:"register,omitempty"`
}
