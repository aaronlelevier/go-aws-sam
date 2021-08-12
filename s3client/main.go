package main

import (
	"flag"
	"fmt"
	"os"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Copy of script here with slight changes:
// https://github.com/aws/aws-sdk-go#complete-sdk-example
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var bucket, key string

	// set CL flags
	flag.StringVar(&bucket, "b", "", "Bucket name.")
	flag.StringVar(&key, "k", "", "Key name.")
	flag.Parse()

	sess := session.Must(session.NewSession(&aws.Config{
			Region: aws.String("us-east-2"),
	}))

	svc := s3.New(sess)

	log.Printf("bucket: %v", bucket)
	log.Printf("key: %v", key)

	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   os.Stdin,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)

}
