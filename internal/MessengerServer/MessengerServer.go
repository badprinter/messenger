package MessengerServer

import (
	"bufio"
	"github.com/badprinter/messenger/internal/config"
	"log"
	"net"
)

type ManagerServer struct {
	listener net.Listener
	Users    *Lobby
}

func NewMessengerServer(cfg config.BaseConfig) (*ManagerServer, error) {
	lis, err := net.Listen("tcp", cfg.Net.Host+":"+cfg.Net.Port)
	if err != nil {
		return nil, err
	}
	return &ManagerServer{
		lis,
		NewLobby(),
	}, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *ManagerServer) Start() {
	m.acceptConnect() // TODO переписать на горутины
}

func (m *ManagerServer) acceptConnect() {
	for { // TODO !m.Quit()
		conn, err := m.listener.Accept()
		if err != nil {
			log.Println(err)
		} else {
			m.Users.Add(conn)
			go m.getMessenge(conn) // TODO рутина
		}
	}
}

func (m *ManagerServer) getMessenge(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // TODO !m.Quit()
		if scanner.Err() != nil {
			log.Println(scanner.Err())
		} else {
			m.Send(conn, scanner.Text())
		}
	}
}

func (m *ManagerServer) Send(sandler net.Conn, msg string) {
	m.Users.Broadcast(sandler, msg)
}
