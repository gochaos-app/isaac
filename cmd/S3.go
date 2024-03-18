package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Save2S3(name string) string {
	varCfg := GetAwsCfg()
	if name == "" {
		name = "prompts.jsonl"
	}
	fileJsonl, err := os.Open(name)
	if err != nil {
		return err.Error()
	}

	defer fileJsonl.Close()
	bucket := varCfg.InitConfig.S3Bucket
	svc := s3.NewFromConfig(*varCfg.AwsConfig)
	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(name),
		Body:   fileJsonl,
	})
	if err != nil {
		fmt.Println("Error uploading to S3", err)
		return err.Error()
	}
	fmt.Println("Successfully uploaded to S3 " + bucket)
	return "Successfully uploaded to S3 " + bucket

}
