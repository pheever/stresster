package main

import (
	"fmt"
	"reflect"

	"gopkg.in/yaml.v3"
)

type foo struct {
	Foo       interface{} `yaml:"foo,omitempty"`
	Bar       int         `yaml:"bar,omitempty"`
	Foobarbaz string      `yaml:"foo.bar.baz,omitempty"`
}

func main() {
	var out foo
	yaml.Unmarshal([]byte(`
foo: 34 % 4
bar: "5"
foo.bar.baz: phivos is the way`), &out)
	fmt.Println(out.Foobarbaz)
	t := reflect.TypeOf(out.Foobarbaz)

	fmt.Println(t)
	fmt.Println(out.Foo)
	t = reflect.TypeOf(out.Bar)
	fmt.Println(t)
	fmt.Println(out.Bar)
}
