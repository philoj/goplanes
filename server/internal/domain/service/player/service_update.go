package playersvc

import "fmt"

func (p *Service) Update(msg []byte) error {
	select {
	case p.msg <- msg:
	default:
		close(p.msg)
		return fmt.Errorf("failed to update game")
	}
	return nil
}
