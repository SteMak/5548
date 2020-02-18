package vanilla

import (
	"github.com/SteMak/vanilla/messages"
	"github.com/SteMak/vanilla/mux"
)

func (bot *module) ping(ctx *mux.Context) error {
	data := map[string]interface{}{
		"Mention": ctx.Message.Author.Mention(),
	}

	m, err := messages.Get("main.ping", data)
	if err != nil {
		return err
	}

	_, err = bot.session.ChannelMessageSendComplex(ctx.Message.ChannelID, m)
	if err != nil {
		return err
	}

	return nil
}

func (bot *module) initCommands() {
	bot.app.Commands = []mux.Command{
		{
			Name:   "ping",
			Action: bot.ping,
		},
	}
}
