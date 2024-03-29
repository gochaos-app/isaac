package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gochaos-app/isaac/cfgisaac"
	"github.com/gochaos-app/isaac/ops"
)

type PromptFn func(string) string

var entries []cfgisaac.FileDB

func switchCommand(IsaacCmd, cmdStr string) string {
	cmdMap := map[string]PromptFn{
		"command":  cmdPromptFn,
		"file":     filePromptFn,
		"document": textFromFileFn,
		"save":     savePrompts,
		"uploadS3": Save2S3,
		"image":    imagePromptFn,
	}
	if _, cmdExists := cmdMap[IsaacCmd]; cmdExists {
		return cmdMap[IsaacCmd](cmdStr)
	} else {
		fmt.Println("Command not found")
		return "Command not found"
	}
}

func ChatBedrock() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("@Isaac → ")
		cmdStr, _ := reader.ReadString('\n')

		cmdStr = strings.TrimSuffix(cmdStr, "\n")

		sysCmds := ops.FindSys(cmdStr)
		//Check if user want to exit isaac
		if len(sysCmds) > 2 {
			if sysCmds[2] == "exit" || sysCmds[2] == "quit" {
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		}
		//Start normal flow
		arrayCmd := ops.FindInst(cmdStr)
		if len(arrayCmd) == 1 {
			//send input to bedrock
			response := ChatBD(cmdStr)
			entries = append(entries, cfgisaac.FileDB{Prompt: cmdStr, Completion: response})

		} else if len(arrayCmd) > 1 {
			//CHeck if a special command was used other than exit and send the input to bedrock
			response := switchCommand(arrayCmd[0], arrayCmd[1])
			entries = append(entries, cfgisaac.FileDB{Prompt: cmdStr, Completion: response})
			// If the input is command ask if the user wants to execute the command
			// If the user types yes, the command will be executed
			if arrayCmd[0] == "command" {
				cmdArray := ops.ExtractCodeBlocks(response)
				if len(cmdArray) == 0 {
					fmt.Println("No command to execute, please try another prompt...")
					return
				}
				cmd := strings.Join(cmdArray, " ")
				fmt.Printf("Execute command? %s ", cmd)
				var userInput string
				fmt.Print("Only yes is accepted: ")
				fmt.Scanf("%s", &userInput)
				if userInput == "yes" {
					executeCommand(cmd)
				} else {
					fmt.Println("Command not executed")
				}
			}

		} else {
			return
		}

	}

}
