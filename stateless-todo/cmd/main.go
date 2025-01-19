package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type Task struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

const ReadHeaderTimeout = 5 * time.Second

func main() {
	tasks := make([]Task, 0)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, _ *http.Request) {
		t, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "marshall error", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(t)
		if err != nil {
			slog.Error("write error", slog.String("error", err.Error()))
		}
	})

	mux.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "invalid json", http.StatusUnprocessableEntity)
			return
		}
		defer r.Body.Close()
		task.Id = len(tasks) + 1
		tasks = append(tasks, task)
	})

	httpServer := http.Server{
		Addr:              "0.0.0.0:5000",
		Handler:           mux,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		slog.Error("server error", slog.String("error", err.Error()))
	}
}
