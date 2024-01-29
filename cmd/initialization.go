package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type AWSConfig struct {
	Region      string `json:"region"`
	Model       string `json:"model"`
	S3Bucket    string `json:"s3bucket"`
	Tokens      string `json:"tokens"`
	Temperature string `json:"temperature"`
}

func FileInit() {

	homeEnv := os.Getenv("HOME")
	filePath := homeEnv + "/.isaac_config.json"
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("This command will ask you for information to customize Isaac")
	fmt.Println("Do you want to continue? ")
	fmt.Println("yes for custom config, anything else for default config")
	response, _ := reader.ReadString('\n')

	if response == "yes\n" {

		_, error := os.Stat(filePath)
		if error == nil {
			fmt.Println("Config file already exists")
			return
		} else {
			fmt.Println("Config file does not exist, creating one")
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Enter AWS region:")
			region, _ := reader.ReadString('\n')
			fmt.Println("Enter S3 bucket name:")
			s3bucket, _ := reader.ReadString('\n')

			fmt.Println("Enter AWS model:")
			model, _ := reader.ReadString('\n')

			fmt.Println("Enter Max tokens:")
			tokens, _ := reader.ReadString('\n')

			fmt.Println("Enter Temperature:")
			temperature, _ := reader.ReadString('\n')

			config := AWSConfig{
				Region:      region[:len(region)-1],
				S3Bucket:    s3bucket[:len(s3bucket)-1],
				Model:       model[:len(model)-1],
				Tokens:      tokens[:len(tokens)-1],
				Temperature: temperature[:len(temperature)-1],
			}
			json.NewEncoder(file).Encode(config)
		}

	} else {
		fmt.Println("Using defaults...")
		config := AWSConfig{
			Region:      "us-east-1",
			S3Bucket:    "isaac-bucket",
			Model:       "ai21.j2-ultra-v1",
			Tokens:      "200",
			Temperature: "0.5",
		}
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("File created")
		json.NewEncoder(file).Encode(config)
	}
}
