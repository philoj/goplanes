package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/philoj/goplanes/server/internal/app/handler"
	"github.com/philoj/goplanes/server/internal/domain/service/lobby"
)

const port = ":8080"

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		log.Print("origin: ", r.Header.Get("origin"))
		return r.Header.Get("origin") == "http://localhost:8081" // FIXME better origin value
	},
}

func main() {
	// Start the lobby
	l := lobbysvc.New()
	go l.Run()

	// Lobby handler websocket endpoint
	lobbyHandler := handler.NewLobbyHandler(upgrader, l)
	http.HandleFunc("/lobby/", lobbyHandler.Handle)

	slog.Info("Starting server", "port", port)
	err := http.ListenAndServe(port, nil)

	// Server exit
	if err != nil {
		log.Fatal("Server failure: ", err)
	}
}
