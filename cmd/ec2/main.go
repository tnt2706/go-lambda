package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
)

var ec2Client *ec2.Client

type Event struct {
	InstanceID string `json:"instance_id"`
	Action     string `json:"action"`
}

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}
	ec2Client = ec2.NewFromConfig(cfg)
}

func handler(ctx context.Context, event Event) {
	if event.Action == "STOP" {
		_, err := ec2Client.StopInstances(ctx, &ec2.StopInstancesInput{
			InstanceIds: []string{event.InstanceID},
			Force:       aws.Bool(false),
		})

		if err != nil {
			log.Fatalf("Unable stop instances: %v", err)
			return
		}
	} else {
		_, err := ec2Client.StartInstances(ctx, &ec2.StartInstancesInput{
			InstanceIds: []string{event.InstanceID},
		})

		if err != nil {
			return
		}
	}

}

func main() {
	lambda.Start(handler)
}
