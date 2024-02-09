package cmd

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Save2S3(name string) string {
	_, _, _, bucket, cfg := GetAwsCfg()
	if name == "" {
		name = "prompts.jsonl"
	}
	fileJsonl, err := os.Open(name)
	if err != nil {
		return err.Error()
	}

	defer fileJsonl.Close()
	svc := s3.NewFromConfig(cfg)
	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(name),
		Body:   fileJsonl,
	})
	if err != nil {
		return err.Error()
	}
	return "Successfully uploaded to S3 " + bucket

}
