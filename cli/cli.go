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
				Name:    "chat",
				Aliases: []string{"c"},
				Usage:   "Chat with Isaac",
				Action: func(c *cli.Context) error {
					cmd.ChatGo()
					return nil
				},
			},
			{
				Name:    "prompt",
				Aliases: []string{"p"},
				Usage:   "make a simple prompt, prompt should be enclosed in quotes",
				Action: func(c *cli.Context) error {
					cfg := cmd.GetAwsCfg()
					cmd.ChatBD(c.Args().First(), cfg)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
