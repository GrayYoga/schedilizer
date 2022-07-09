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
	Num           int           `json:"-"`
}

type TaskList struct {
	Tasks []Task
}

type TasksDurations struct {
	Durations      time.Duration
	HumanDurations string
}

var (
	runned bool
	wgSync sync.WaitGroup
)

func (t *TaskList) Add(duration string, mode string, number int) Task {
	timeDuration, _ := time.ParseDuration(duration)
	task := Task{timeDuration, duration, mode, number}
	t.Tasks = append(t.Tasks, task)
	return task
}

func (t *TaskList) Pop() Task {
	first := t.Tasks[0]
	t.Tasks = t.Tasks[1:]
	return first
}

func (t *TaskList) GetList() (string, error) {
	if len(t.Tasks) > 1 {
		data, err := json.Marshal(t.Tasks[1:])
		if err != nil {
			return "", err
		} else {
			return string(data), nil
		}
	} else {
		return "[]", nil
	}
}

func (t *TaskList) GetDurations() (string, error) {
	var sum time.Duration
	if len(t.Tasks) > 1 {
		for _, task := range t.Tasks[1:] {
			sum = sum + task.Duration
		}
	}

	data, err := json.Marshal(TasksDurations{
		Durations:      sum,
		HumanDurations: fmt.Sprintf("%v", sum),
	})
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}

func (t *TaskList) RunSerial(wait bool) {
	if runned {
		log.Printf("Queued %d tasks. Current run schedulled", len(t.Tasks))
		if wait {
			wgSync.Wait()
		}
		return
	} else {
		runned = true
		wgSync.Add(1)
		for len(t.Tasks) > 0 {

			locT := t.Tasks[0]
			locT.doTask()
			t.Pop()
		}
		wgSync.Done()
		runned = false
	}
}

func (t *Task) doTask() {
	log.Printf("Start task #%d duration: %v\n", t.Num, t.Duration)
	time.Sleep(t.Duration)
	log.Printf("Stop task  #%d after %v\n", t.Num, t.Duration)
}
