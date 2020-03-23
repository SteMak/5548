package main

import (
	"os"

	"github.com/SteMak/vanilla/out"

	_ "github.com/SteMak/vanilla/modules/dropper"
	_ "github.com/SteMak/vanilla/modules/vanilla"
)

func main() {
	if err := app().Run(os.Args); err != nil {
		out.Fatal(err)
	}

}
