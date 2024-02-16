package cmd

import (
	"encoding/json"
	"os"
)

func savePrompts(name string) string {
	if name == "" {
		name = "prompts.jsonl"
	}
	tmpfile, err := os.Create(name)
	if err != nil {
		return err.Error()
	}
	defer tmpfile.Close()
	for _, d := range entries {
		jsonData, err := json.Marshal(d)
		if err != nil {
			return err.Error()
		}
		tmpfile.WriteString(string(jsonData) + "\n")
	}
	return "Prompts saved"
}
