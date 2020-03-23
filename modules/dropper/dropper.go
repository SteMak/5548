package dropper

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cam-per/discordgo"
)

type module struct {
	session *discordgo.Session
	config  config

	running bool
}

func (module) ID() string {
	return "dropper"
}

func (bot *module) IsRunning() bool {
	return bot.running
}

func (bot *module) Init(prefix, configPath string) error {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &bot.config)
	if err != nil {
		return err
	}
	return nil
}

func (bot *module) Start(session *discordgo.Session) {
	bot.session = session
	bot.running = true
}

func (bot *module) Stop() {
	bot.running = false
}
