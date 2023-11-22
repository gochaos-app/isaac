package cmd

import (
	"encoding/json"
	"log"
	"os"
)

type fileDB struct {
	Prompt   string `json:"prompt"`
	Response string `json:"response"`
}

func savePrompts(entries []fileDB) {
	tmpfile, err := os.Create("prompts.json")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()

	encoder := json.NewEncoder(tmpfile)
	err = encoder.Encode(entries)
	if err != nil {
		log.Fatal(err)
	}
}
