package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/philoj/goplanes/server/internal/app/api/handler/lobby/create"
	"github.com/philoj/goplanes/server/internal/app/api/handler/lobby/join"
	"github.com/philoj/goplanes/server/internal/app/socket"
	"github.com/philoj/goplanes/server/internal/domain/model"
)

type Dependencies struct {
	AuthSvc  AuthService
	LobbySvc createlobby.LobbyService
}

type Config struct {
	Port int `mapstructure:"port"`
}

type AuthService interface {
	Middleware(next http.Handler) http.Handler

	GetUser(ctx context.Context) (model.User, error)
}

func Start(ctx context.Context, dep Dependencies, cfg Config) error {
	protectedRouter := http.NewServeMux()
	protectedRouter.Handle("POST /lobbies", createlobby.NewHandler(dep.AuthSvc, dep.LobbySvc))
	protectedRouter.Handle("GET /lobbies/{lobbyID}/join", joinlobby.NewHandler(socket.NewUpgrader()))

	rootRouter := http.NewServeMux()
	rootRouter.Handle("/", dep.AuthSvc.Middleware(protectedRouter))

	slog.InfoContext(ctx, "Starting API server", "port", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), rootRouter)
}
