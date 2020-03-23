package main

import (
	"fmt"
	"os"

	"github.com/SteMak/vanilla/out"

	_ "github.com/SteMak/vanilla/modules/dropper"
	_ "github.com/SteMak/vanilla/modules/vanilla"
)

func main() {
	out.Infoln("Environment:")
	for _, env := range os.Environ() {
		out.Infoln(env)
	}

	fmt.Println()

	if err := app().Run(os.Args); err != nil {
		out.Fatal(err)
	}
}
