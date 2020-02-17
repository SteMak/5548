package modules

import (
	"github.com/bwmarrin/discordgo"
)

type Module interface {
	LoadConfig(string) error

	Name() string

	Start(prefix string, session *discordgo.Session)
	Stop()
}

var (
	modules  = make(map[string]Module)
	attached = make([]Module, 0)
)

func Register(name string, module Module) {
	modules[name] = module
}

func Get(name string) Module {
	if module, ok := modules[name]; ok {
		return module
	}
	return nil
}

func Attach(module Module) {
	attached = append(attached, module)
}
