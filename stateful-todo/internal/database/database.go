package database

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	user     = "******"
	password = "******"
	host     = "*postgres" // "192.168.2.51"
	port     = "*5432"     // "30000"
	db       = "*stateful-todo-db"
)

// MustGetPgxPool get pgxpool or os.Exit(1)
func MustGetPgxPool(ctx context.Context) *pgxpool.Pool {
	log := slog.With("func", "database.MustGetPgxPool")

	dbpool, err1 := NewPgxPool(ctx)
	if err1 != nil {
		log.Warn("Failed init postgres", slog.String("error", err1.Error()))
		os.Exit(1)
	}

	return dbpool
}

func NewPgxPool(ctx context.Context) (*pgxpool.Pool, error) {
	log := slog.With("func", "database.NewPgxPool")

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	db = os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v",
		host,
		port,
		user,
		password,
		db,
	)

	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("database.NewPgxPool: %w", err)
	}

	dbpool, err1 := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err1 != nil {
		log.Warn("Error connecting to the database", slog.String("error", err1.Error()))
		return nil, fmt.Errorf("database.NewPgxPool: %w", err1)
	}

	err2 := dbpool.Ping(ctx) // эта команда заменяет acquire + ping
	if err2 != nil {
		log.Warn("Could not ping database", slog.String("error", err2.Error()))
		return nil, fmt.Errorf("database.NewPgxPool: %w", err2)
	}

	log.Info("successfully connected to database")
	return dbpool, nil
}
