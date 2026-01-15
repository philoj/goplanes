package lobbysvc

func (l *Lobby) PlayerExists(id int) bool {
	_, exists := l.players[id]
	return exists
}
