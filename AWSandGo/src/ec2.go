package main

// START OMIT
import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/ec2"
	"os"
)

var region = "us-west-2"

func main() {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	e := ec2.New(creds, region, nil)
	desc, err := e.DescribeIntances(nil) // gets an instance of DescribeInstancesResults
	if err != nil {
		fmt.Fatalln(err)
	}
	for i, r := range desc.Reservations {
		fmt.Printf("%2d: %v - %v - %v\n", i+1,
			*r.Instances[i].InstanceID, *r.Instances[i].Tags[0].Value, *r.Instances[i].State.Name,
		)
	}
}

// END OMIT
