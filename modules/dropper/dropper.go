package dropper

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
)

type module struct {
	session *discordgo.Session
	config  config

	prefix string
}

func (module) ID() string {
	return "dropper"
}

func (m *module) LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &m.config)
	if err != nil {
		return err
	}
	return nil
}

func (m *module) Start(prefix string, session *discordgo.Session) {
	m.session = session
	m.prefix = prefix
}

func (m *module) Stop() {

}
