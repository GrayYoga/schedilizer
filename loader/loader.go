package loader

import (
	"bufio"
	"os"
	"schedulizer/types"
	"time"
)

func LoadTasks(fileName string) ([]types.Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return tasks, nil
}
