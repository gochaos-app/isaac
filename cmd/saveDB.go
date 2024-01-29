package cmd

import (
	"encoding/json"
	"log"
	"os"
)

type fileDB struct {
	Prompt     string `json:"prompt"`
	Completion string `json:"completion"`
}

func savePrompts(entries []fileDB) {
	tmpfile, err := os.Create("prompts.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()
	for _, d := range entries {
		jsonData, err := json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}
		tmpfile.WriteString(string(jsonData) + "\n")
	}

}
