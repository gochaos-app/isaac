package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gochaos-app/isaac/ops"
)

func ChatGo(config *AWSConfig) {
	region := config.Region
	reader := bufio.NewReader(os.Stdin)
	cfg := GetAwsCfg(region)
	var entries []fileDB
	for {
		fmt.Print("@Isaac: ")
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdStr = strings.TrimSuffix(cmdStr, "\n")

		if len(ops.FindSys(cmdStr)) > 2 {
			subArray := ops.FindSys(cmdStr)
			if subArray[1] == "sys." {
				switch subArray[2] {
				//   Special commands for isaac
				case "exit":
					fmt.Println("Goodbye!")
					os.Exit(0)
				case "save":
					fmt.Println("Saving...")
					savePrompts(entries)
				case "":
					fmt.Println("No command found")
				default:
					fmt.Println("Command not found")
				}
			}
		} else if len(ops.FindCmdIgnoreParams(cmdStr)) > 1 {
			cmdSlice := strings.Fields(ops.FindCmdIgnoreParams(cmdStr)[1])
			cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			fmt.Println(cmd.Run())

		} else {
			response := ChatBD(cmdStr, config.Model, config.Tokens, config.Temperature, cfg)
			entries = append(entries, fileDB{Prompt: cmdStr, Response: response})
		}
	}
}
