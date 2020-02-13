package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	vanilla "github.com/SteMak/vanilla"
	"github.com/SteMak/vanilla/config"
	"github.com/SteMak/vanilla/out"
	"github.com/urfave/cli"
)

func run(c *cli.Context) error {
	fmt.Println("Bot is running. Press Ctrl + C to exit.")

	config.Load(c.GlobalString("config"))
	out.SetDebug(c.GlobalBool("debug"))

	vanilla.Run()
	defer vanilla.Stop()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	return nil
}
