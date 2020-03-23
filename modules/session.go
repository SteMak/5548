package modules

import (
	"github.com/SteMak/vanilla/config"
	"github.com/SteMak/vanilla/out"
	"github.com/cam-per/discordgo"
)

var (
	session *discordgo.Session
)

func authentificate() {
	out.Infoln("\nAuthentifications...")
	s, err := discordgo.New(config.Session.Token)
	if err != nil {
		out.Fatal(err)
	}
	session = s

	session.StateEnabled = true

	session.SyncEvents = false

	session.AddHandler(onReady)

	if err := session.Open(); err != nil {
		out.Fatal(err)
	}
	out.Infoln("websocket started")

	out.Infoln("authorized as:", session.State.User.String())
	out.Debugln("token:", s.Token)
}
