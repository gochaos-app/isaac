package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func ReadInitFile() *AWSConfig {
	var config AWSConfig
	homeEnv := os.Getenv("HOME")
	filePath := homeEnv + "/.isaac_config.json"
	_, error := os.Stat(filePath)
	if os.IsNotExist(error) {
		fmt.Println("Config file doesnt exists")
		fmt.Println("Please run isaac init")
		os.Exit(1)
	}
	// Open our jsonFile
	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

func GetAwsCfg() (string, string, string, string, aws.Config) {

	isaacCfg := ReadInitFile()

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(isaacCfg.Region))
	if err != nil {
		fmt.Println("Unable to load SDK config, " + err.Error())
		log.Fatal(err)
	}

	return isaacCfg.Model, isaacCfg.Tokens, isaacCfg.Temperature, isaacCfg.S3Bucket, cfg
}
