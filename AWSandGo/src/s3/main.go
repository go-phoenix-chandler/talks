package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/s3"
	"log"
	"os"
)

var region = "us-west-2"
var bucket = "gomeetup"

// Connect
func Connect() *s3.S3 {
	creds := aws.Creds(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
	return s3.New(creds, region, nil)
}

// Upload
func Upload(s *s3.S3) {
	resp, err := s.PutObject(&s3.PutObjectRequest{
		Body:,
		Bucket: bucket,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

// Download
func Download(s *s3.S3) {
	resp, err := s.PutObject(&s3.PutObjectRequest{})
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println("")
	os.Exit(0)
}
