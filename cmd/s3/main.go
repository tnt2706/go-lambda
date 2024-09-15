package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func handler(ctx context.Context, s3Event events.S3Event) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("Error loading default configuration: %v. Have you set up your AWS account?", err)
		return
	}

	s3Client := s3.NewFromConfig(sdkConfig)

	for _, record := range s3Event.Records {
		s3Record := record.S3

		// Copy the S3 object
		copySource := fmt.Sprintf("%s/%s", s3Record.Bucket.Name, s3Record.Object.Key)
		result, err := s3Client.CopyObject(ctx, &s3.CopyObjectInput{
			Bucket:     aws.String("tinhtn-lambda-2"), // Target bucket
			CopySource: aws.String(copySource),        // Source object (bucket/key)
			Key:        aws.String(s3Record.Object.Key),
		})

		if err != nil {
			log.Printf("Failed to copy object from %s to bucket %s: %v", copySource, "tinhtn-lambda-2", err)
			continue
		}

		log.Printf("Successfully copied object: %v", result)
		log.Printf("[%s - %s] Bucket = %s, Key = %s", record.EventSource, record.EventTime, s3Record.Bucket.Name, s3Record.Object.Key)
	}
}

func main() {
	// Start the AWS Lambda handler
	lambda.Start(handler)
}
