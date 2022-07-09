package server

import (
	"fmt"
	"net/http"
)

func Run(addr string) error {

	http.HandleFunc("/add", add)
	http.HandleFunc("/schedule", schedule)
	http.HandleFunc("/time", time)

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
		_, err := fmt.Fprintf(w, "Mode is a %s, duration is %s\n", mode, duration)
		if err != nil {
			return
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
		// TODO тут массив актуальных задач, стоящих в очереди на выполнение, в формате JSON.
		_, err := fmt.Fprintf(w, "[]")
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

func time(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// TODO оставшееся время на выполнение всех находящихся в очереди задач.
		_, err := fmt.Fprintf(w, "queue is empty")
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
