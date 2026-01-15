package playersvc

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/philoj/goplanes/server/internal/domain/model"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// msg pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Service struct {
	id    int
	lobby model.LobbyService

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	msg chan []byte
}

func New(id int, l model.LobbyService, conn *websocket.Conn) model.PlayerService {
	return &Service{id: id, lobby: l, conn: conn, msg: make(chan []byte, 256)}
}
