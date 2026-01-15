package lobbysvc

import "github.com/philoj/goplanes/server/internal/domain/model"

func (l *Lobby) JoinLobby(p model.PlayerService) {
	l.join <- p
}
