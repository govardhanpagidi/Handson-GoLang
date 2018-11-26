package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)
//**********How to run**************
// Lists all objects in a bucket
//
// Usage:
// listObjects <bucket>
func main() {
	if len(os.Args) < 2 {
		fmt.Println("you must specify a bucket")
		return
	}
	bucket := os.Args[1]
	fmt.Println("bucket name : ",bucket)

	
	conf := aws.Config{Region: aws.String("us-west-2")}
	sess := session.New(&conf)

	svc := s3.New(sess)

	resp, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		fmt.Println("Unable to list items in bucket \n", bucket, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("")
		fmt.Println("Name         :", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size (MB)    :", *item.Size/1000)
		fmt.Println("Storage class:", *item.StorageClass)
	}

}
