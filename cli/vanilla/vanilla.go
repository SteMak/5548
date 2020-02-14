package main

import "github.com/urfave/cli"

func app() *cli.App {
	a := &cli.App{
		Name:        "Picker",
		Description: "Auto-picking sakura bot for Anilibria discord server",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "d, debug",
				Usage: "Enable debug mode",
			},
			&cli.StringFlag{
				Name:  "c, config",
				Value: "config.json",
				Usage: "Path to the config file",
			},
		},
		Action: run,
		Commands: []cli.Command{
			{
				Name:   "migrate",
				Usage:  "Create database",
				Action: migrate,
			},
		},
	}
	a.UseShortOptionHandling = true
	return a
}
