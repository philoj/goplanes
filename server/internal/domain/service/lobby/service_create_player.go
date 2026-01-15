package lobbysvc

import (
	"github.com/gorilla/websocket"
	"github.com/philoj/goplanes/server/internal/domain/service/player"
)

func (l *Lobby) CreatePlayer(id int, conn *websocket.Conn) {
	p := playersvc.New(id, l, conn)
	l.JoinLobby(p)
	p.Run()
}
