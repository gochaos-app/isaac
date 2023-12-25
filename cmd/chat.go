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
				case "cmd":
					cmdArray := subArray[3:]
					cmdSlice := strings.Fields(cmdArray[0])
					if len(cmdSlice) == 0 {
						fmt.Println("No command found")
					} else {
						cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						fmt.Println(cmd.Run())
					}
				case "":
					fmt.Println("Command not found")
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

		} else if len(ops.FindLoadIgnoreParams(cmdStr)) > 1 {
			fileStr := ops.FindLoadIgnoreParams(cmdStr)[1]
			txtFile, err := ops.LoadFile(fileStr)
			if err != nil {
				fmt.Println("Error reading file")
				return
			}
			txtFile = "make a summary of the following text: " + txtFile

			response := ChatBD(txtFile, config.Model, config.Tokens, config.Temperature, cfg)
			entries = append(entries, fileDB{Prompt: txtFile, Response: response})
		} else if len(ops.MakeSummaryIgnoreParams(cmdStr)) > 1 {
			fileStr := ops.MakeSummaryIgnoreParams(cmdStr)[1]
			txtFile, err := ops.LoadFile(fileStr)
			if err != nil {
				fmt.Println("Error reading file")
				return
			}
			txtFile = "Make a summary of the following text: " + txtFile

			response := ChatBD(txtFile, config.Model, config.Tokens, config.Temperature, cfg)
			entries = append(entries, fileDB{Prompt: txtFile, Response: response})
		} else {
			response := ChatBD(cmdStr, config.Model, config.Tokens, config.Temperature, cfg)
			entries = append(entries, fileDB{Prompt: cmdStr, Response: response})
		}
	}
}
