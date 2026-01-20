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
)

const port = ":8080"

func main() {
	ctx := context.Background()
	mux := http.NewServeMux()

	lobbySvc := lobbysvc.NewService(nil)

	mux.Handle("POST /lobbies", createlobby.NewHandler(lobbySvc))
	mux.Handle("GET /lobbies/{lobbyID}/join", joinlobby.NewHandler(socket.NewUpgrader()))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.ErrorContext(ctx, "Server exited with error", err)
		os.Exit(1)
	}
}
