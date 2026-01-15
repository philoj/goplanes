package lobbysvc

import (
	"github.com/philoj/goplanes/server/internal/domain/model"
)

type SocketPayload struct {
	Id  int
	Msg []byte
}

// Lobby the game lobby for players to join
type Lobby struct {
	// Registered clients.
	players map[int]model.PlayerService

	//  msg messages to the clients.
	msg chan SocketPayload

	// join requests from the clients.
	join chan model.PlayerService

	// leave requests from clients.
	leave chan model.PlayerService
}

func New() *Lobby {
	return &Lobby{
		msg:     make(chan SocketPayload),
		join:    make(chan model.PlayerService),
		leave:   make(chan model.PlayerService),
		players: make(map[int]model.PlayerService),
	}
}
