package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Save2S3(cfg aws.Config, bucket string) {

	fileJsonl, err := os.Open("prompts.jsonl")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileJsonl.Close()
	svc := s3.NewFromConfig(cfg)
	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("prompts.jsonl"),
		Body:   fileJsonl,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Successfully uploaded to S3", bucket)

}
