package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/badprinter/messenger/internal/MessengerServer"
	"github.com/badprinter/messenger/internal/config"
	"log"
	"net"
)

var (
	configPath = flag.String("config", "settings.toml", "path to config file")
)

var conns []net.Conn

func main() {
	flag.Parse()
	var cfg config.BaseConfig
	_, err := toml.DecodeFile(*configPath, &cfg)
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	server, err := MessengerServer.NewMessengerServer(cfg)
	if err != nil {
		log.Printf("Error creating server: %v", err)
	}

	server.Start()

}
