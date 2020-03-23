package modules

import (
	"os"
	"path/filepath"

	"github.com/SteMak/vanilla/messages"

	"github.com/SteMak/vanilla/config"
	"github.com/SteMak/vanilla/out"
)

func loadTemplates() {
	out.Infoln("Loading templates...")
	err := filepath.Walk(config.Bot.Templates, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		err = messages.AddTpl(path)
		if err != nil {
			out.Fatal(err)
		}

		out.Infoln(path)
		return nil
	})

	if err != nil {
		out.Fatal(err)
	}
}
