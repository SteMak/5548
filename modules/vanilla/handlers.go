package vanilla

import (
	"github.com/SteMak/vanilla/util"
	"github.com/cam-per/discordgo"
)

func (bot *module) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !util.EqualAny(m.ChannelID, bot.config.Channels) {
		return
	}

	if m.Content == "" {
		return
	}

	bot.app.Run(m.Message, m.Content)
}
