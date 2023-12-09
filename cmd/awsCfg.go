package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func GetAwsCfg(region string) aws.Config {

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Unable to load SDK config, " + err.Error())
		log.Fatal(err)
	}
	return cfg
}
