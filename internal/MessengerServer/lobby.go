package MessengerServer

import (
	"github.com/badprinter/messenger/internal/user"
	"sync"
)

type Lobby struct {
	sync.Mutex
	array []user.User // map
}

func NewLobby() *Lobby {
	return &Lobby{}
}

func (l *Lobby) Add(User *user.User) {
	l.Lock()
	defer l.Unlock()
	l.array = append(l.array, *User)
}

func (l *Lobby) RemoveConnection(User user.User) {
	l.Lock()
	defer l.Unlock()
	for i, u := range l.array {
		if u.GetIP() == User.GetIP() {
			l.array = append(l.array[:i], l.array[i+1:]...)
			u.CloseTunnel()
			return
		}
	}
}

func (l *Lobby) CloseAll() {
	l.Lock()
	defer l.Unlock()
	for _, u := range l.array {
		u.CloseTunnel()
	}
	l.array = nil
}

func (l *Lobby) Broadcast(who user.User, msg string) {
	for _, u := range l.array {
		if u.GetIP() != who.GetIP() {
			u.Say(msg + "\n")
		}
	}
}
