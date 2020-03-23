package vanilla

import (
	"encoding/json"
	"io/ioutil"

	"github.com/SteMak/vanilla/router"

	"github.com/cam-per/discordgo"
)

type module struct {
	session *discordgo.Session
	config  config

	app *router.App

	running bool
}

func (module) ID() string {
	return "vanilla"
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

	bot.app = &router.App{
		Prefix:      prefix,
		Description: bot.config.Description,
	}

	bot.initCommands()

	return nil
}

func (bot *module) Start(session *discordgo.Session) {
	bot.session = session
	bot.running = true

	bot.session.AddHandler(bot.onMessageCreate)
}

func (bot *module) Stop() {
	bot.running = false
}
