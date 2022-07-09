package main

import (
	"log"
	"schedulizer/loader"
	"schedulizer/runner"
)

func main() {
	tasks, err := loader.LoadTasks("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	runner.Run(tasks)
}
