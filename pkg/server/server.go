package server

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/badprinter/messenger/internal/config"
)

type Messenger struct {
	cfg          *config.NetCofnig
	service      net.Listener
	stoperChan   chan bool
	MessengeChan chan string
}

func NewMessenger(cfg *config.BaseConfig) *Messenger {
	return &Messenger{
		&cfg.Net,
		nil,
		make(chan bool),
		make(chan string),
	}
}

func (m *Messenger) Run() error {
	var err error
	m.service, err = net.Listen("tcp", m.cfg.Host+":"+m.cfg.Port)
	if err != nil {
		return err
	}
	go stop(m.service, m.stoperChan)
	go readMesseng(m.service, m.MessengeChan)

	return nil
}

func readMesseng(l net.Listener, ch chan string) {
	c, err := l.Accept()
	if err != nil {
		log.Fatalln("str:45 ", err)
	}
	log.Println("Connect is accept.")
	reader := bufio.NewReader(c)
	for {
		messange, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		ch <- messange
	}
}

func stop(web net.Listener, stopChan chan bool) {
	defer web.Close()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-stopChan:
			log.Println("Close port by exit.")
			return
		case <-sigChan:
			log.Println("Close port by signal.")
			return
		}
	}
}

func (m *Messenger) Stop() {
	m.stoperChan <- true
}
