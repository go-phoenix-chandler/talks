package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/route53"
	"log"
	"os"
)

var region = "us-west-1"

// Connect will create a valid ec2 client
func Connect() *route53.Route53 {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "") // HL
	return route53.New(creds, region, nil)
}

// ListZones prints out the hosted zones
func ListZones(r *route53.Route53) {
	resp, err := r.ListHostedZones(&route53.ListHostedZonesRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*resp)
}

//
func GetStatus(r *route53.Route53) {
	resp, err := r.GetHealthCheckStatus( // HL
		&route53.GetHealthCheckStatusRequest{ // HL
			HealthCheckID: aws.String("gomeetup_check"), // HL
		}, // HL
	) // HL
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*resp)
}

func main() {
	r := Connect()
	//ListZones(r)
	GetStatus(r)
}
