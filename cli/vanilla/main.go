package main

import (
	"os"

	"github.com/SteMak/vanilla/out"
)

func main() {
	if err := app().Run(os.Args); err != nil {
		out.Fatal(err)
	}
}
