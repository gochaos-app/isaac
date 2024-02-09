package cli

import (
	"log"
	"os"

	"github.com/gochaos-app/isaac/cmd"
	"github.com/urfave/cli/v2"
)

func InitCli() {
	app := &cli.App{
		Name:  "Isaac",
		Usage: " CLI that can help you on various tasks",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialize Isaac",
				Action: func(c *cli.Context) error {

					cmd.FileInit()
					return nil
				},
			},
			{
				Name:    "chat",
				Aliases: []string{"c"},
				Usage:   "Chat with Isaac",
				Action: func(c *cli.Context) error {
					cmd.ChatBedrock()
					return nil
				},
			},
			{
				Name:    "prompt",
				Aliases: []string{"p"},
				Usage:   "make a simple prompt, prompt should be enclosed in quotes",
				Action: func(c *cli.Context) error {

					cmd.ChatBD(c.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
