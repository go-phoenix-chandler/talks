package main

// START OMIT
import (
	"fmt"
	"log"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/rds"
)

var region = "us-west-2"

func main() {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	db := rds.New(creds, region, nil)
	desc, err := db.DescribeDBInstances(nil)
	if err != nil {
		log.Fatalln(err)
	}
	for i, d := range desc.DBInstances {
		fmt.Printf("%2d: %v\n", i+1, *d.DBInstanceIdentifier)
	}
}

// END OMIT
