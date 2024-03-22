package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gochaos-app/isaac/ops"
)

const commandPrompt = "You are an expert in Linux Operating Systems and the task is to provide bash commands. Return ONLY a command inside a code Block, for example ```ls```. DO NOT USE single or double quotes, or any other character. Give me ONLY the command ready to run in a BASH terminal. The command should do: "

func cmdPromptFn(cmdStr string) string {
	complete := commandPrompt + cmdStr
	response := ChatBD(complete)
	return response
}

func imagePromptFn(cmdStr string) string {
	response := ImageBD(cmdStr)
	return response
}

func textFromFileFn(cmdStr string) string {
	SliceFile := strings.Fields(cmdStr)
	filename := SliceFile[0]

	prompt := strings.Join(SliceFile[1:], " ")
	TextFromDoc(filename, prompt)
	return ""
}

func filePromptFn(cmdStr string) string {
	SliceFile := strings.Fields(cmdStr)
	filename := SliceFile[0]
	prompt := strings.Join(SliceFile[1:], " ")

	file, err := ops.LoadFile(filename)

	if err != nil {
		fmt.Println("Error reading file")
		return ""
	}
	complete := prompt + ": " + file
	response := ChatBD(complete)
	return response
}

func executeCommand(cleanCmd string) {

	CmdSlice := strings.Fields(cleanCmd)
	cmd := exec.Command(CmdSlice[0], CmdSlice[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.Run())

}
