package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/sqs"
	"log"
	"os"
)

var region = "us-west-2"
var queue = "gomeetup"

// Connect will create a valid ec2 client
func Connect() *sqs.SQS {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	return sqs.New(creds, region, nil)
}

func GetQueues(s *sqs.SQS) []string {
	resp, err := s.ListQueues(&sqs.ListQueuesRequest{
		QueueNamePrefix: aws.String(queue),
	})
	if err != nil {
		log.Fatalln(err)
	}
	return resp.QueueURLs
}

func main() {
	s := Connect()
	for i, q := range GetQueues(s) {
		fmt.Printf("%2d: %s\n", i+1, q)
	}
}
