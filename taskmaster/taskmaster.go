package taskmaster

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type Task struct {
	Duration      time.Duration `json:"duration"`
	HumanDuration string        `json:"human_duration"`
	Mode          string        `json:"Mode"`
	Num           int           `json:"Number"`
}

type TaskList struct {
	Tasks []Task
}

func (t *TaskList) Add(duration string, mode string, number int) {
	timeDuration, _ := time.ParseDuration(duration)
	t.Tasks = append(t.Tasks, Task{timeDuration, duration, mode, number})
}

func (t *TaskList) Next() Task {
	first := t.Tasks[0]
	t.Tasks = t.Tasks[1:]
	return first
}

func (t *TaskList) GetList() (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func (t *TaskList) GetDurations() string {
	var sum time.Duration
	for _, task := range t.Tasks {
		sum = sum + task.Duration
	}
	return fmt.Sprintf("%s", sum)

}

var wg sync.WaitGroup

func (t *TaskList) RunSerial() {
	t.LimitedRun(1)
}

func (t *TaskList) LimitedRun(limit int) {
	ch := make(chan int, limit)
	for len(t.Tasks) > 0 {
		wg.Add(1)
		locT := t.Next()
		ch <- 1
		go locT.limitedTask(ch)
	}
	wg.Wait()
}

func (t *Task) limitedTask(ch chan int) {
	defer func() {
		_ = <-ch
		defer wg.Done()
	}()
	log.Printf("Start task #%d duration: %v\n", t.Num, t.Duration)
	time.Sleep(t.Duration)
	log.Printf("Stop task  #%d after %v\n", t.Num, t.Duration)
}
