package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/ec2"
	"log"
	"os"
)

var region = "us-west-2"

// Connect will create a valid ec2 client
func Connect() *ec2.EC2 {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "") // HL
	return ec2.New(creds, region, nil)
}

// Stop will shutdown an EC2 instance
func Stop(e *ec2.EC2, instances []string) {
	resp, err := e.StopInstances(&ec2.StopInstancesRequest{ // HL
		DryRun:      aws.Boolean(false), // HL
		Force:       aws.Boolean(true),  // HL
		InstanceIDs: instances,          // HL
	}) // HL
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*resp)
}

// Start will bring up an EC2 instance
func Start(e *ec2.EC2, instances []string) {
	resp, err := e.StartInstances(&ec2.StartInstancesRequest{ // HL
		AdditionalInfo: aws.String(""),     // HL
		DryRun:         aws.Boolean(false), // HL
		InstanceIDs:    instances,          // HL
	}) // HL
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*resp)
}

// ListInstanceData will print out instance data
func ListInstanceData(e *ec2.EC2) {
	resp, err := e.DescribeInstances(nil) // HL
	if err != nil {
		log.Fatalln(err)
	}
	for i, r := range resp.Reservations {
		fmt.Printf("%2d: %v - %v - %v\n", i+1,
			*r.Instances[i].InstanceID, *r.Instances[i].Tags[0].Value, *r.Instances[i].State.Name, // HL
		)
	}
}
