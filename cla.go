package main

import (
	"flag"
)

var (
	planPath    string
	planThreads int
	planRampUp  string
)

//CLALoad the command line arguments
func CLALoad() {
	flag.StringVar(&planPath, "plan", "./plan.yaml", "path to a plan yaml file")
	flag.IntVar(&planThreads, "threads", 1, "how many threads to run, overwrites plan configured threads")
	flag.StringVar(&planRampUp, "rampup", "0", "plan ramp up time, overwrites plan configured threads")
	flag.Parse()
}
