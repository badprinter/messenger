package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/badprinter/messenger/internal/config"
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

	var number uint32 = 0
	for {
		select {
		case msg, ok := <-m.MessengeChan:
			if ok {
				fmt.Printf("%3d: %s\n", number, msg)
			}
		default:
			continue
		}
	}

}
