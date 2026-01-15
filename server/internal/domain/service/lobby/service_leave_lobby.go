package lobbysvc

import "github.com/philoj/goplanes/server/internal/domain/model"

func (l *Lobby) LeaveLobby(p model.PlayerService) {
	l.leave <- p
}
