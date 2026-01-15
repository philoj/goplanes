package playersvc

func (p *Service) JoinLobby() {
	p.lobby.JoinLobby(p)
}
