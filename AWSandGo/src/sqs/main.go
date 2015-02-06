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
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "") // HL
	return sqs.New(creds, region, nil)
}

// SendMessage will get a message from the given queue
func SendMessage(s *sqs.SQS, msg, q string) {
	_, err := s.SendMessage( // HL
		&sqs.SendMessageRequest{ // HL
			DelaySeconds: aws.Integer(0),  // HL
			MessageBody:  aws.String(msg), // HL
			QueueURL:     aws.String(q),   // HL
		}, // HL
	) // HL
	if err != nil {
		log.Fatalln(err)
	}
}

// RetrieveMessage will get a message from the given queue
func RetrieveMessage(s *sqs.SQS, q string) []sqs.Message {
	resp, err := s.ReceiveMessage( // HL
		&sqs.ReceiveMessageRequest{ // HL
			QueueURL: aws.String(q), // HL
		}, // HL
	) // HL
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Messages
}

// GetQueues returns a string slice of all found queues by their URL
func GetQueues(s *sqs.SQS) []string {
	resp, err := s.ListQueues( // HL
		&sqs.ListQueuesRequest{ // HL
			QueueNamePrefix: aws.String(queue), // HL
		}, // HL
	) // HL
	if err != nil {
		log.Fatalln(err)
	}
	//for i, q := range GetQueues(s) {
	//	fmt.Printf("%2d: %s\n", i+1, q)
	//}
	return resp.QueueURLs
}

// PrintMessages will do what you think...
func PrintMessages(msgs []sqs.Message) {
	for i, m := range msgs {
		fmt.Printf("%2d: %s: %s\n", i+1, *m.MessageID, *m.Body)
	}
}

func main() {
	s := Connect()
	SendMessage(s, "Test Message 1", GetQueues(s)[0])
	time.Sleep(1 * time.Second)
	PrintMessages(RetrieveMessage(s, GetQueues(s)[0]))
}
