package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gochaos-app/isaac/cmd"
	"github.com/urfave/cli/v2"
)

var config cmd.AWSConfig

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
					isaacCfg := readFile()
					cmd.ChatGo(isaacCfg)
					return nil
				},
			},
			{
				Name:    "prompt",
				Aliases: []string{"p"},
				Usage:   "make a simple prompt, prompt should be enclosed in quotes",
				Action: func(c *cli.Context) error {
					isaacCfg := readFile()
					cfg := cmd.GetAwsCfg(isaacCfg.Region)
					cmd.ChatBD(c.Args().First(), isaacCfg.Model, isaacCfg.Tokens, isaacCfg.Temperature, cfg)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func readFile() *cmd.AWSConfig {
	homeEnv := os.Getenv("HOME")
	filePath := homeEnv + "/.isaac_config.json"
	_, error := os.Stat(filePath)
	if os.IsNotExist(error) {
		fmt.Println("Config file doesnt exists")
		fmt.Println("Please run isaac init")
		os.Exit(1)
	}
	// Open our jsonFile
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened config file")
	return &config
}
