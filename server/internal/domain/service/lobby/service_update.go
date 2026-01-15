package lobbysvc

func (l *Lobby) Update(id int, msg []byte) {
	l.msg <- SocketPayload{id, msg}
}
