package model

import "github.com/gorilla/websocket"

type PlayerService interface {
	Id() int
	JoinLobby()
	LeaveLobby()
	Run()
	Update(msg []byte) error
}

type LobbyService interface {
	Run()
	JoinLobby(p PlayerService)
	LeaveLobby(p PlayerService)
	Update(id int, msg []byte)
	CreatePlayer(id int, conn *websocket.Conn)
}
