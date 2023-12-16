package cmd

import (
	"fmt"
	"os"
	"strings"
)

func LoadFile(file string) (string, error) {
	var content string
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		fmt.Println("File does not exists")
		return "", err
	} else {
		fmt.Println("File does exist")
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file")
			return "", err
		}
		content = string(data)
		content = strings.ReplaceAll(content, "\n", " ")
	}
	return content, nil

}
