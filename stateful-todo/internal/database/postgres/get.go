package postgres

import (
	"context"

	"github.com/arslanovdi/k8s_examples/stateful-todo/internal/model"
	"github.com/jackc/pgx/v5"
)

func (r *Repo) GetTasks() ([]model.Task, error) {
	// сборка запроса - query
	query, args, err1 := psql.Select("*").
		From("tasks").
		ToSql()

	if err1 != nil {
		return nil, err1
	}

	rows, _ := r.dbpool.Query(context.Background(), query, args...) //nolint:errcheck checked further in pgx.CollectRows
	defer rows.Close()

	tasks, err2 := pgx.CollectRows(rows, pgx.RowToStructByName[model.Task])
	if err2 != nil {
		return nil, err2
	}

	return tasks, nil
}
