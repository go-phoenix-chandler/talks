package main

// START OMIT
import (
	"fmt"
	"os"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/ec2"
)

var region = "us-west-2"

func main() {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	e := ec2.New(creds, region, nil)
	desc, err := e.DescribeIntances(nil) // gets an instance of DescribeInstancesResults
	if err != nil {
		fmt.Fatalln(err)
	}
	// interate over the []Reservation slice and print them
	for i, r := range desc.Reservations.Instances {
		fmt.Printf("%2d: %10s - %10s\n", i+1, r.InstanceID, r.State.Name)
	}
}

// END OMIT
