package cli

import (
	"errors"
	"fmt"
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
				Name:    "document",
				Aliases: []string{"d"},
				Usage:   "Get text out of a document, file or an image",
				Action: func(c *cli.Context) error {
					filename := c.Args().Get(0)
					if _, err := os.Stat(filename); err != nil {
						err := errors.New("File does not exist")
						return err
					}
					cmd.TextFromDoc(filename, "")
					return nil
				},
			},
			{
				Name:    "image",
				Aliases: []string{"img"},
				Usage:   "Make an image with a prompt",
				Action: func(c *cli.Context) error {
					cmd.ImageBD(c.Args().First())
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
				Usage:   "Make a simple prompt, prompt should be enclosed in quotes",
				Action: func(c *cli.Context) error {

					fmt.Println(cmd.ChatBD(c.Args().First()))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
