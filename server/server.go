package server

import (
	"fmt"
	"net/http"
	"schedulizer/taskmaster"
)

var taskList taskmaster.TaskList
var number int

func Run(addr string) error {
	http.HandleFunc("/add", add)
	http.HandleFunc("/schedule", schedule)
	http.HandleFunc("/time", getTime)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			_, err := fmt.Fprintf(w, "ParseForm() err: %v", err)
			if err != nil {
				return
			}
			return
		}
		mode := r.FormValue("mode")
		duration := r.FormValue("duration")
		number = number + 1
		taskList.Add(duration, mode, number)
		switch mode {
		case "sync":
			taskList.RunSerial()
			_, err := fmt.Fprintf(w, "Mode is a %s, duration is %s\n", mode, duration)
			if err != nil {
				return
			}
		case "async":
			go taskList.RunSerial()
			_, err := fmt.Fprintf(w, "Mode is a %s, duration is %s\n", mode, duration)
			if err != nil {
				return
			}
		default:
			_, err := fmt.Fprintf(w, "Sorry, only (sync, async) values in field `mode` are supported.")
			if err != nil {
				return
			}
		}

	default:
		_, err := fmt.Fprintf(w, "Sorry, only POST method are supported.")
		if err != nil {
			return
		}
	}
}

func schedule(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		data, err := taskList.GetList()
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(w, data)
		if err != nil {
			return
		}
	default:
		_, err := fmt.Fprintf(w, "Sorry, only GET method are supported.")
		if err != nil {
			return
		}
	}
}

func getTime(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		_, err := fmt.Fprintf(w, taskList.GetDurations())
		if err != nil {
			return
		}
	default:
		_, err := fmt.Fprintf(w, "Sorry, only GET method are supported.")
		if err != nil {
			return
		}
	}
}
