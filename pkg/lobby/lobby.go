package lobby

import (
	"net"
)

type Lobby struct {
	tunels []net.Conn
}

func NewLobby() *Lobby {
	return &Lobby{}
}

func (l *Lobby) Add(tunel net.Conn) {
	l.tunels = append(l.tunels, tunel)
}

func (l *Lobby) GetTunels() []net.Conn {
	return l.tunels
}
