package playersvc

func (p *Service) LeaveLobby() {
	p.lobby.LeaveLobby(p)
	close(p.msg)
}
