package modules

import (
	"runtime/debug"
	"time"

	"github.com/SteMak/vanilla"

	"github.com/SteMak/vanilla/config"
	"github.com/SteMak/vanilla/messages"
	"github.com/SteMak/vanilla/out"
)

func Send(channelID string, tplName string, data interface{}) (err error) {
	m, err := messages.Get(tplName, data)
	if err != nil {
		out.Err(true, err)
		return
	}

	_, err = session.ChannelMessageSendComplex(channelID, m)
	if err != nil {
		out.Err(true, err)
		return
	}
	return
}

func SendError(msg string) {
	data := map[string]interface{}{
		"Timestamp": time.Now().UTC().Format(time.StampNano),
		"Version":   vanilla.Vesion,
		"Message":   msg,
		"Stack":     string(debug.Stack()),
	}

	m, err := messages.Get("main/error.xml", data)
	if err != nil {
		out.Err(false, err)
		return
	}

	_, err = session.ChannelMessageSendComplex(*config.Bot.ErrorsChannel, m)
	if err != nil {
		out.Err(false, err)
	}
}
