package main

import (
	"log"
	"schedulizer/args"
	"schedulizer/loader"
	"schedulizer/runner"
)

func main() {
	a := args.Args{}
	a.ParseArgs()
	tasks, err := loader.LoadTasks("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found tasks: %d, limit: %d", len(tasks), a.Limit)
	runner.LimitedRun(tasks, a.Limit)
}
