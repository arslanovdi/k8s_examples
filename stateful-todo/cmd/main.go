package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/arslanovdi/k8s_examples/stateful-todo/internal/api"
	"github.com/arslanovdi/k8s_examples/stateful-todo/internal/database"
	"github.com/arslanovdi/k8s_examples/stateful-todo/internal/database/postgres"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	dbpool := database.MustGetPgxPool(context.Background())

	slog.Info("Migration started")
	if err := goose.Up(stdlib.OpenDBFromPool(dbpool), // получаем соединение с базой данных из пула
		"migrations"); err != nil {
		slog.Warn("Migration failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	repo := postgres.NewPostgresRepo(dbpool)

	todoServer := api.NewServer(repo, "0.0.0.0:5000")

	todoServer.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	slog.Info("Graceful shutdown")
	todoServer.Stop()
	dbpool.Close()
}
