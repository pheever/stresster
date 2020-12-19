package main

import (
	"fmt"

	"github.com/pheever/stresster/plan"
)

func main() {
	CLALoad()
	pln, er := plan.Load(planPath)
	if er != nil {
		fmt.Printf("failed: %v", er)
		return
	}
	fmt.Printf("the plan:\n%+v\n", pln)
}
