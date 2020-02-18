package vanilla

import (
	"encoding/json"
	"io/ioutil"

	"github.com/SteMak/vanilla/router"

	"github.com/bwmarrin/discordgo"
)

type module struct {
	session *discordgo.Session
	config  config

	app *router.App
}

func (module) ID() string {
	return "vanilla"
}

func (bot *module) LoadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &bot.config)
	if err != nil {
		return err
	}
	return nil
}

func (bot *module) Start(prefix string, session *discordgo.Session) {
	bot.session = session

	bot.app = &router.App{
		Prefix:      prefix,
		Description: bot.config.Description,
		Session:     session,
	}

	bot.initCommands()

	bot.session.AddHandler(bot.onMessageCreate)
}

func (bot *module) Stop() {

}
