package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"schedulizer/taskmaster"
)

var taskList taskmaster.TaskList
var number int

func Run(addr string) error {
	var router = mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/add", add)
	router.HandleFunc("/schedule", schedule)
	router.HandleFunc("/time", getTime)
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	return http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
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
		added := taskList.Add(duration, mode, number)
		switch mode {
		case "sync":
			taskList.RunSerial(true)
			err := json.NewEncoder(w).Encode(added)
			if err != nil {
				return
			}
		case "async":
			go taskList.RunSerial(false)
			err := json.NewEncoder(w).Encode(added)
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
		_, err = fmt.Fprint(w, data)
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
		dur, err := taskList.GetDurations()
		if err != nil {
			return
		}
		_, err = fmt.Fprintf(w, dur)
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
