package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gochaos-app/isaac/ops"
)

type PromptFn func(string) string

var entries []fileDB

func switchCommand(IsaacCmd, cmdStr string) string {
	cmdMap := map[string]PromptFn{
		"command":    cmdPromptFn,
		"kubernetes": k8sPromptFn,
		"file":       filePromptFn,
		"save":       savePrompts,
		"uploadS3":   Save2S3,
		"image":      imagePromptFn,
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
			entries = append(entries, fileDB{Prompt: cmdStr, Completion: response})

		} else if len(arrayCmd) > 1 {
			//CHeck if a special command was used other than exit and send the input to bedrock
			response := switchCommand(arrayCmd[0], arrayCmd[1])
			entries = append(entries, fileDB{Prompt: cmdStr, Completion: response})
			// If the input is command ask if the user wants to execute the command
			// If the user types yes, the command will be executed
			if arrayCmd[0] == "command" {
				cmd := ops.CleanCmd(response)
				fmt.Println("Execute command? Only yes is accepted:" + cmd)
				var userInput string
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
