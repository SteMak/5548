package vanilla

import (
	"github.com/SteMak/vanilla/out"

	"github.com/SteMak/vanilla/config"

	"github.com/bwmarrin/discordgo"
)

var (
	bot vanilla
)

type vanilla struct {
	session *discordgo.Session
}

func authorize() {
	s, err := discordgo.New(config.Session.Token)
	if err != nil {
		out.Fatal(err)
	}

	if err := s.Open(); err != nil {
		out.Fatal(err)
	}
	out.Infoln("websocket started")

	out.Infoln("authorized as:", s.State.User.String())
	out.Debugln("token:", s.Token)

	s.SyncEvents = false
	bot.session = s
}

func Run() {
	authorize()
}

func Stop() {
	bot.session.Close()
}
