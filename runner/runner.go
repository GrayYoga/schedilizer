package runner

import (
	"log"
	"schedulizer/types"
	"time"
)

func Run(tasks []types.Task) {
	for _, t := range tasks {
		task(t)
	}
}
func task(task types.Task) {
	log.Printf("Start task #%d duration: %v\n", task.Num, task.Dur)
	time.Sleep(task.Dur)
	log.Printf("Stop task  #%d\n", task.Num)
}
