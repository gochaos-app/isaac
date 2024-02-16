package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gochaos-app/isaac/cfgisaac"
)

func ReadInitFile() *cfgisaac.InitConfig {
	var cfg cfgisaac.InitConfig
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
	err = json.Unmarshal(jsonFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func GetAwsCfg() cfgisaac.IsaacConfig {

	isaacCfg := ReadInitFile()

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(isaacCfg.Region))
	if err != nil {
		fmt.Println("Unable to load SDK config, " + err.Error())
		log.Fatal(err)
	}
	isaac := cfgisaac.IsaacConfig{
		InitConfig: *isaacCfg,
		AwsConfig:  &cfg,
	}

	return isaac
}
