package main

// START OMIT
import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/s3"
	"log"
	"os"
)

var region = "us-west-2"

func main() {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	creds := aws.Creds(accessKey, secretKey, "")
	s := s3.New(creds, region, nil)

	os.Exit(0)
}

// END OMIT
