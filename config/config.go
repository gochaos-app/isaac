package config

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type IsaacConfig struct {
	InitConfig InitConfig  `json:"initConfig"`
	AwsConfig  *aws.Config `json:"awsConfig"`
}

type InitConfig struct {
	Region      string `json:"region"`
	Model       string `json:"model"`
	ImageModel  string `json:"imageModel"`
	S3Bucket    string `json:"s3bucket"`
	Tokens      string `json:"tokens"`
	Temperature string `json:"temperature"`
}

type FileDB struct {
	Prompt     string `json:"prompt"`
	Completion string `json:"completion"`
}
