package api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/arslanovdi/k8s_examples/stateful-todo/internal/model"
)

const ReadHeaderTimeout = 5 * time.Second

type Repo interface {
	GetTasks() ([]model.Task, error)
	AddTask(model.Task) (int, error)
}

type TaskAPI struct {
	db     Repo
	mux    *http.ServeMux
	server http.Server
}

func NewServer(db Repo, addr string) *TaskAPI {
	mux := http.NewServeMux()

	api := TaskAPI{
		db:  db,
		mux: mux,
	}

	api.addRoutes()

	api.server = http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}
	return &api
}

func (t *TaskAPI) Run() {
	go func() {
		slog.Info("starting http server", slog.String("address", t.server.Addr))
		if err := t.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", slog.String("error", err.Error()))
		}
	}()
}

func (t *TaskAPI) Stop() {
	if err := t.server.Shutdown(context.Background()); err != nil {
		slog.Error("server shutdown error", slog.String("error", err.Error()))
	}
}

func (t *TaskAPI) addRoutes() {
	t.mux.HandleFunc("GET /tasks", t.handlerGetTasks)
	t.mux.HandleFunc("POST /tasks", t.handlerAddTask)
}

func (t *TaskAPI) handlerGetTasks(w http.ResponseWriter, _ *http.Request) {
	tasks, err := t.db.GetTasks()
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "marshall error", http.StatusInternalServerError)
		return
	}
	_, err = w.Write(result)
	if err != nil {
		slog.Error("write error", slog.String("error", err.Error()))
	}
}

func (t *TaskAPI) handlerAddTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "invalid json", http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	id, err := t.db.AddTask(task)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}
	slog.Debug("task added", slog.Int("id", id))
}
