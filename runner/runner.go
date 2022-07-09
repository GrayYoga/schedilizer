package runner

import (
	"log"
	"schedulizer/types"
	"sync"
	"time"
)

var wg sync.WaitGroup

func LimitedRun(tasks []types.Task, limit int) {
	ch := make(chan int, limit)
	for _, t := range tasks {
		wg.Add(1)
		locT := t
		ch <- 1
		go limitedTask(locT, ch)
	}
	wg.Wait()
}

func limitedTask(task types.Task, ch chan int) {
	defer func() {
		_ = <-ch
		defer wg.Done()
	}()
	log.Printf("Start task #%d duration: %v\n", task.Num, task.Dur)
	time.Sleep(task.Dur)
	log.Printf("Stop task  #%d after %v\n", task.Num, task.Dur)
}
