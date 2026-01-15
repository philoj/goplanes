package lobbysvc

import "log"

func (l *Lobby) Run() {
	for {
		select {
		case p := <-l.join:
			log.Printf("player %d join", p.Id())
			l.players[p.Id()] = p
		case client := <-l.leave:
			if _, ok := l.players[client.Id()]; ok {
				delete(l.players, client.Id())
			}
		case b := <-l.msg:
			for id := range l.players {
				// do not write back to the same player
				if b.Id == id {
					continue
				}
				//log.Println("Sending to id:", p.playerId, string(b.msg))
				err := l.players[id].Update(b.Msg)
				if err != nil {
					delete(l.players, id)
				}
			}
		}
	}
}
