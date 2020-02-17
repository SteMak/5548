package vanilla

import (
	"github.com/SteMak/vanilla/config"
	"github.com/SteMak/vanilla/messages"
	"github.com/SteMak/vanilla/out"
	"github.com/bwmarrin/discordgo"
)

func (bot *vanilla) onReady(s *discordgo.Session, e *discordgo.Ready) {
	data := map[string]interface{}{
		"Name": e.User,
	}

	m, err := messages.Get("main.started", data)
	if err != nil {
		out.Err(err)
		return
	}

	if config.Bot.LogChannel != nil {
		_, err := bot.session.ChannelMessageSendComplex(*config.Bot.LogChannel, m)
		if err != nil {
			out.Err(err)
		}
	}
}
