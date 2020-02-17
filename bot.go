package vanilla

import (
	"github.com/SteMak/vanilla/config"
	"github.com/SteMak/vanilla/messages"
	"github.com/SteMak/vanilla/modules"
	"github.com/SteMak/vanilla/out"
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
	bot.session = s

	initHandlers()

	if err := s.Open(); err != nil {
		out.Fatal(err)
	}
	out.Infoln("websocket started")

	out.Infoln("authorized as:", s.State.User.String())
	out.Debugln("token:", s.Token)

	s.SyncEvents = false
}

func loadMessages() {
	if err := messages.LoadMessage(config.Bot.Messages); err != nil {
		out.Err(err)
	}
}

func loadModules() {
	out.Infoln("\nLoading modules...")
	if len(*config.Modules) == 0 {
		out.Fatal("No attached modules. Check the config file")
	}

	for name, module := range *config.Modules {
		out.Infof("  %-25s", name)

		m := modules.Get(name)
		if m == nil {
			out.Infoln("[NOT IMPLEMENTED]")
			continue
		}

		if !module.Enabled {
			if m == nil {
				out.Infoln("[OFF]")
				continue
			}
		}

		if err := m.LoadConfig(module.Config); err != nil {
			out.Err(err)
			continue
		}

		m.Start(module.Prefix, bot.session)
		modules.Attach(m)

		if err := messages.LoadMessage(module.Messages); err != nil {
			out.Err(err)
			continue
		}

		out.Infoln("[ON]")
	}
}

func initHandlers() {
	bot.session.AddHandler(bot.onReady)
}

func Run() {
	loadMessages()
	authorize()
	loadModules()
}

func Stop() {
	bot.session.Close()
}
