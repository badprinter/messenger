package server

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/badprinter/messenger/internal/config"
	"github.com/badprinter/messenger/pkg/lobby"
)

type Messenger struct {
	cfg          *config.NetCofnig
	service      net.Listener
	lobby        *lobby.Lobby
	quitChan     chan bool // Когда в канал будет записано True, все горутины должны завершится
	MessengeChan chan string
}

func NewMessenger(cfg *config.BaseConfig) *Messenger {
	return &Messenger{
		&cfg.Net,
		nil,
		lobby.NewLobby(),
		make(chan bool),
		make(chan string),
	}
}

func (m *Messenger) IsQuit() bool {
	select {
	case res := <-m.quitChan:
		return res
	default:
		return false
	}
}

// TODO вынести
func (m *Messenger) DoCommand(cmd string) {
	switch cmd {
	case "/quit":
		m.QuitMessanger()
	case "/help":
		println("Print help messange.")
	}
}

func (m *Messenger) Run() error {
	var err error
	m.service, err = net.Listen("tcp", m.cfg.Host+":"+m.cfg.Port)
	if err != nil {
		return err
	}
	go m.lobbyManager()
	go m.exitHandler()

	return nil
}

func (m *Messenger) SendMessenge(msg string) {
	tunnels := m.lobby.GetTunels()
	for _, v := range tunnels {
		go v.Write([]byte(msg))
	}
}

func (m *Messenger) lobbyManager() {
	for !m.IsQuit() {
		c, err := m.service.Accept()
		if err != nil {

			log.Println("lobbyManager ", err)
			continue
		}
		m.lobby.Add(c)
		log.Printf("Connect: %s is accept.", c.RemoteAddr().String())
		go m.readMesseng(c, m.MessengeChan)
	}
}

func (m *Messenger) readMesseng(c net.Conn, ch chan string) {
	scanner := bufio.NewScanner(c)
	for scanner.Scan() || !m.IsQuit() {
		if scanner.Err() != nil {
			return
		}
		ch <- scanner.Text()
	}
}

func (m *Messenger) GetMesseng() string {
	msg, ok := <-m.MessengeChan
	if ok {
		return msg
	}
	return ""
}

// TODO переписать входные параметры
func (m *Messenger) exitHandler() {
	defer func() {
		tunels := m.lobby.GetTunels()
		if len(tunels) != 0 {
			for _, v := range tunels {
				v.Close()
			}
		}
		m.service.Close()
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-m.quitChan:
		log.Println("Close port by exit.")
	case <-sigChan:
		m.QuitMessanger()
		log.Println("Close port by signal.")
	}
}

func (m *Messenger) QuitMessanger() {
	m.quitChan <- true
}

func test_function(b bool) {
}
