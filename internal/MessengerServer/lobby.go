package MessengerServer

import (
	"net"
	"sync"
)

type Lobby struct {
	sync.Mutex
	array []net.Conn
}

func NewLobby() *Lobby {
	return &Lobby{}
}

func (l *Lobby) Add(conn net.Conn) {
	l.Lock()
	defer l.Unlock()
	l.array = append(l.array, conn)
}

func (l *Lobby) RemoveConnection(conn net.Conn) {
	l.Lock()
	defer l.Unlock()
	for i, c := range l.array {
		if c == conn {
			l.array = append(l.array[:i], l.array[i+1:]...)
			return
		}
	}
}

func (l *Lobby) CloseAll() {
	l.Lock()
	defer l.Unlock()
	for _, conn := range l.array {
		conn.Close()
	}
	l.array = nil
}

func (l *Lobby) Broadcast(who net.Conn, msg string) {
	send := append([]byte(msg), '\n')
	for _, conn := range l.array {
		if conn.RemoteAddr().String() != who.RemoteAddr().String() {
			conn.Write(send)
		}
	}
}
