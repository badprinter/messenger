package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/badprinter/messenger/internal/config"
	"github.com/badprinter/messenger/internal/inputcontrol"
	"github.com/badprinter/messenger/pkg/server"
)

var (
	configPath = flag.String("c", "settings.toml", "Path to settings file.")
)

func main() {
	flag.Parse()
	var cfg config.BaseConfig
	_, err := toml.DecodeFile(*configPath, &cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	m := server.NewMessenger(&cfg)
	err = m.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Server run with options:\nHost: %s\nPort: %s\n\n", cfg.Net.Host, cfg.Net.Port)

	go PrintMessengs(m)
	SendMsg(m)
}

func SendMsg(m *server.Messenger) {
	var msg string
	log.Print("Sender is runed.")
	for !m.IsQuit() {
		res, err := fmt.Scanf("%s", &msg)
		if res > 0 && err == nil {
			if inputcontrol.IsCommand(msg) {
				m.DoCommand(msg)
			} else {
				m.SendMessenge(msg)
			}
		}
	}
}

func PrintMessengs(m *server.Messenger) {
	var msg string
	for {
		msg = m.GetMesseng()
		if msg != "" {
			fmt.Printf("%s\n", msg)
		}
	}
}
