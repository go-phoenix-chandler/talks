package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/cloudwatch"
)

var region = "us-west-2"

// Connect will provide a valid RDS client
func Connect() *cloudwatch.CloudWatch {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "") // HL
	return cloudwatch.New(creds, region, nil)
}

// ListMetrics will ...list metrics
func ListMetrics(c *cloudwatch.CloudWatch) {
	resp, err := c.ListMetrics(&cloudwatch.ListMetricsInput{}) // HL
	if err != nil {
		log.Fatalln(err)
	}
	for i, m := range resp.Metrics {
		fmt.Printf("%3d: %s\n", i+1, *m.MetricName)
	}
}

//

func main() {
	c := Connect()
	ListMetrics(c)
}
