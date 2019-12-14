package main

import (
	"fmt"

	"github.com/pheever/stresster/core"
)

func main() {
	plan, err := core.LoadPlanFromFile("./test-plan2.yml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("loaded plan:\n\t%+v", plan)
}
