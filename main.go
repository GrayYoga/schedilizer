package main

import (
	"bufio"
	"log"
	"os"
	"schedulizer/runner"
	"schedulizer/types"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var tasks []types.Task
	num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		dur, _ := time.ParseDuration(str)
		tasks = append(tasks, types.Task{Num: num, Dur: dur})
		num = num + 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	runner.Run(tasks)
}
