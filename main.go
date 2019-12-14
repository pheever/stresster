package main

import (
	"fmt"

	"github.com/pheever/stresster/core"
)

func main() {
	plan, err := core.LoadPlanFromFile("./examples/test-plan.yml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf(plan.Name)
}
