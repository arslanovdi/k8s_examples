package postgres

import (
	"context"

	"github.com/arslanovdi/k8s_examples/stateful-todo/internal/model"
)

func (r *Repo) AddTask(task model.Task) (int, error) {
	// сборка запроса - query
	query, args, err1 := psql.Insert("tasks").
		Columns("title", "description").
		Values(task.Title, task.Description).
		Suffix("RETURNING id").
		ToSql()

	if err1 != nil {
		return -1, err1
	}

	err2 := r.dbpool.QueryRow(context.Background(), query, args...).Scan(&task.Id)
	if err2 != nil {
		return -1, err2
	}

	return task.Id, nil
}
