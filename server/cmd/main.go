package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/philoj/goplanes/server"
	"github.com/philoj/goplanes/server/internal/app/api"
	"github.com/philoj/goplanes/server/internal/app/configparse"
	lobbyrepo "github.com/philoj/goplanes/server/internal/domain/repository/lobby"
	authsvc "github.com/philoj/goplanes/server/internal/domain/service/auth"
	lobbysvc "github.com/philoj/goplanes/server/internal/domain/service/lobby"
	"github.com/philoj/goplanes/server/internal/infra/postgres"
	"github.com/philoj/goplanes/server/migrations"
)

func main() {
	ctx := context.Background()

	// Parse config
	var cfg server.Config
	err := configparse.Unmarshal(&cfg, configparse.FromPath("config.yaml"))
	if err != nil {
		slog.ErrorContext(ctx, "Failed to parse config", "err", err)
		os.Exit(1)
	}

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

	// Auto-migrate DB
	if cfg.AutoMigrate {
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
	}

	// Auth Service
	authSvc := authsvc.NewService()

	// Lobby Service
	lobbyRepo := lobbyrepo.NewRepository(db)
	lobbySvc := lobbysvc.NewService(lobbyRepo)

	// API Server
	serverDep := api.Dependencies{
		AuthSvc:  authSvc,
		LobbySvc: lobbySvc,
	}
	err = api.Start(ctx, serverDep, cfg.Server)
	if err != nil {
		slog.ErrorContext(ctx, "API Server exited with error", err)
		os.Exit(1)
	}
}
