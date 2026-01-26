package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	createlobby "github.com/philoj/goplanes/server/internal/app/handler/lobby/create"
	"github.com/philoj/goplanes/server/internal/app/handler/lobby/join"
	"github.com/philoj/goplanes/server/internal/app/socket"
	lobbysvc "github.com/philoj/goplanes/server/internal/domain/service/lobby"
	"github.com/philoj/goplanes/server/internal/infra/postgres"
	"github.com/philoj/goplanes/server/migrations"
)

const port = ":8080"

func main() {
	ctx := context.Background()
	mux := http.NewServeMux()

	// GORM database client
	db, err := postgres.NewClient("user=root dbname=goplanes sslmode=disable host=localhost")
	if err != nil {
		slog.ErrorContext(ctx, "Failed to connect to postgres database", "err", err)
		os.Exit(1)
	}
	// The native sql.DB object from the postgres client.
	sqlDB, err := db.DB()
	if err != nil {
		slog.ErrorContext(ctx, "Failed to get sql.DB from postgres client", "err", err)
		os.Exit(1)
	}

	slog.InfoContext(ctx, "Running DB migrations")
	defer func() {
		cErr := sqlDB.Close()
		if cErr != nil {
			slog.ErrorContext(ctx, "Failed to close sql.DB", "err", cErr)
		}
	}()
	if err = migrations.Migrate(ctx, sqlDB); err != nil {
		slog.ErrorContext(ctx, "Failed to run migrations", "err", err)
		os.Exit(1)
	}

	lobbySvc := lobbysvc.NewService(nil)

	mux.Handle("POST /lobbies", createlobby.NewHandler(lobbySvc))
	mux.Handle("GET /lobbies/{lobbyID}/join", joinlobby.NewHandler(socket.NewUpgrader()))

	slog.InfoContext(ctx, "Starting server", "port", port)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.ErrorContext(ctx, "Server exited with error", err)
		os.Exit(1)
	}
}
