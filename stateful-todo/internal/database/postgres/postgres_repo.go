package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // Squirell плэйсхолдер для Postgres

type Repo struct {
	dbpool *pgxpool.Pool
}

func NewPostgresRepo(dbpool *pgxpool.Pool) *Repo {
	return &Repo{
		dbpool: dbpool,
	}
}
