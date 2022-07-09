package runner

import (
	"log"
	"schedulizer/types"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Run(tasks []types.Task) {

	for _, t := range tasks {
		wg.Add(1)
		locT := t
		go task(locT)
	}
	wg.Wait()
}
func task(task types.Task) {
	defer wg.Done()
	log.Printf("Start task #%d duration: %v\n", task.Num, task.Dur)
	time.Sleep(task.Dur)
	log.Printf("Stop task  #%d\n", task.Num)
}
