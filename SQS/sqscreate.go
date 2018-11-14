package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"os"
)

// Source : https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sqs-example-create-queue.html

// Prerequisites
//Create a IAM user in AWS and use credentials to create  a session

//How to run

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please enter a name to create queue..")
		return
	}

	qName := os.Args[1]

	fmt.Println("Start..")

	//Create a session
	conf := aws.Config{Region: aws.String("us-west-2")}
	sess := session.New(&conf)

	//Create a SQS service client
	svc := sqs.New(sess)

	fmt.Println("Creating a Queue..")

	result, err2 := svc.CreateQueue(&sqs.CreateQueueInput{QueueName: aws.String(qName), Attributes: map[string]*string{
		"DelaySeconds":           aws.String("60"),
		"MessageRetentionPeriod": aws.String("86400")},
	})

	if err2 != nil {
		fmt.Println("Error : ", err2)
	}

	fmt.Println("Success", *result.QueueUrl)

	fmt.Println("End.")
}
