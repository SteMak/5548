package vanilla

import (

	"github.com/SteMak/vanilla/config"
	"github.com/bwmarrin/discordgo"
)

type console struct {
	timeFormat string
}

var (
	out = console{
		timeFormat: "15:04:05.000",
	}
	bot picker
)

type picker struct {
	session *discordgo.Session

	debug bool
	watch bool
	mode  int
}

func initSession() *discordgo.Session {
	s, err := discordgo.New(config.Session.Token)
	if err != nil {
		out.fatal(err)
	}

	if err := s.Open(); err != nil {
		out.fatal("cannot open websocket:", err)
	}
	out.infoln("websocket started")

	out.infoln("authorized as:", s.State.User.String())
	out.debugln("token:", s.Token)

	s.SyncEvents = false
	return s
}

func Run() {
	bot.session = initSession()
}

func Stop() {
	bot.session.Close()
	out.infoln("websocket closed")
}