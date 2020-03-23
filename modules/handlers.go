package modules

import (
	"github.com/SteMak/vanilla/config"
	"github.com/cam-per/discordgo"
)

func onReady(s *discordgo.Session, e *discordgo.Ready) {
	if config.Bot.LogChannel == nil {
		return
	}

	data := map[string]interface{}{
		"Name": e.User,
	}

	Send(*config.Bot.LogChannel, "main/started.xml", data)
}
