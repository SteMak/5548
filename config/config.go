package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var cfg config

var (
	Session = &cfg.Session
	Bot     = &cfg.Bot
)

type config struct {
	Session session `json:"session,omitempty"`
	Bot     bot     `json:"bot,omitempty"`
}

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config loaded:", path)
}
