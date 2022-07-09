package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"schedulizer/taskmaster"
	"sync"
	"testing"
	"time"
)

func TestAsyncAddTask(t *testing.T) {
	resp, err := http.PostForm("http://localhost:8090/add", url.Values{
		"mode":     []string{"async"},
		"duration": []string{"1s"},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var task taskmaster.Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		t.Fatalf("Error decode response: %s\n", err)
	}
	assert.Equal(t, taskmaster.Task{
		Duration:      time.Duration(1000000000),
		HumanDuration: "1s",
		Mode:          "async",
	}, task, "Response incorrect")
}

func TestSyncAddTask(t *testing.T) {
	resp, err := http.PostForm("http://localhost:8090/add", url.Values{
		"mode":     []string{"sync"},
		"duration": []string{"2s"},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var task taskmaster.Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		t.Fatalf("Error decode response: %s\n", err)
	}
	assert.Equal(t, task, taskmaster.Task{
		Duration:      time.Duration(2000000000),
		HumanDuration: "2s",
		Mode:          "sync",
		Num:           0,
	}, "Response incorrect")
}

func TestAddTasks(t *testing.T) {
	var wg sync.WaitGroup
	var expected []taskmaster.Task

	for i := 0; i < 5; i++ {
		expected = append(expected, taskmaster.Task{
			Duration:      time.Duration(3000000000),
			HumanDuration: "3s",
			Mode:          "async",
		})
	}
	for _, task := range expected {
		wg.Add(1)
		task := task
		go func() {
			defer wg.Done()
			_, err := http.PostForm("http://localhost:8090/add", url.Values{
				"mode":     []string{task.Mode},
				"duration": []string{task.HumanDuration},
			})
			if err != nil {
				t.Error(err)
				return
			}
		}()
	}
	wg.Wait()

	resp, err := http.Get("http://localhost:8090/schedule")
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var tasks []taskmaster.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		t.Fatalf("Error decode response: %s\n", err)
	}
	assert.Equal(t, expected[1:], tasks, "Response incorrect")

	resp, err = http.Get("http://localhost:8090/time")
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var tasksDur taskmaster.TasksDurations
	err = json.NewDecoder(resp.Body).Decode(&tasksDur)
	if err != nil {
		t.Fatalf("Error decode response: %s\n", err)
	}
	assert.Equal(t, taskmaster.TasksDurations{
		Durations:      time.Duration(12000000000),
		HumanDurations: "12s",
	}, tasksDur, "Response incorrect")

	resp, err = http.PostForm("http://localhost:8090/add", url.Values{
		"mode":     []string{"sync"},
		"duration": []string{"0s"},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var task taskmaster.Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		t.Fatalf("Error decode response: %s\n", err)
	}
	task.Num = 0
	assert.Equal(t, task, taskmaster.Task{
		Duration:      time.Duration(0),
		HumanDuration: "0s",
		Mode:          "sync",
		Num:           0,
	}, "Response incorrect")
}
