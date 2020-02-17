package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/SteMak/vanilla/out"
)

var cfg config

var (
	Session = &cfg.Session
	Bot     = &cfg.Bot
	Storage = &cfg.Storage
	Modules = &cfg.Modules
)

type config struct {
	Session session           `json:"session,omitempty"`
	Bot     bot               `json:"bot,omitempty"`
	Storage storage           `json:"storage,omitempty"`
	Modules map[string]module `json:"modules,omitempty"`
}

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		out.Fatal(err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		out.Fatal(err)
	}
	out.Infoln("Config loaded:", path)
}
