package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/sqs"
	"log"
	"os"
	"time"
)

var region = "us-west-2"
var queue = "gomeetup"

// Connect will create a valid ec2 client
func Connect() *sqs.SQS {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	return sqs.New(creds, region, nil)
}

// SendMessage will get a message from the given queue
func SendMessage(s *sqs.SQS, msg, q string) {
	_, err := s.SendMessage(
		&sqs.SendMessageRequest{
			DelaySeconds: aws.Integer(0),
			MessageBody:  aws.String(msg),
			QueueURL:     aws.String(q),
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
}

// RetrieveMessage will get a message from the given queue
func RetrieveMessage(s *sqs.SQS, q string) []sqs.Message {
	resp, err := s.ReceiveMessage(
		&sqs.ReceiveMessageRequest{
			QueueURL: aws.String(q),
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Messages
}

// PrintMessages will do what you think...
func PrintMessages(msgs []sqs.Message) {
	for i, m := range msgs {
		fmt.Printf("%2d: %s: %s\n", i+1, *m.MessageID, *m.Body)
	}
}

// GetQueues returns a string slice of all found queues by their URL
func GetQueues(s *sqs.SQS) []string {
	resp, err := s.ListQueues(
		&sqs.ListQueuesRequest{
			QueueNamePrefix: aws.String(queue),
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	return resp.QueueURLs
}

func main() {
	s := Connect()
	/*
		for i, q := range GetQueues(s) {
			fmt.Printf("%2d: %s\n", i+1, q)
		}
	*/
	SendMessage(s, "Test Message 1", GetQueues(s)[0])
	time.Sleep(1 * time.Second)
	//RetrieveMessage(s, GetQueues(s)[0])
	PrintMessages(RetrieveMessage(s, GetQueues(s)[0]))
}
