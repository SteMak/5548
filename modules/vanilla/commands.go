package vanilla

import (
	"github.com/SteMak/vanilla/modules"
	"github.com/SteMak/vanilla/router"
)

func (bot *module) ping(ctx *router.Context) error {
	data := map[string]interface{}{
		"Mention": ctx.Message.Author.Mention(),
	}

	modules.Send(ctx.Message.ChannelID, "main/ping.xml", data)
	return nil
}

func (bot *module) initCommands() {
	bot.app.Commands = []router.Command{
		{
			Name:   "ping",
			Action: bot.ping,
		},
	}
}
