package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Source : https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sqs-example-create-queue.html

// Prerequisites
//Create a IAM user in AWS and use credentials to create  a session

//How to run
func main() {

	// if len(os.Args) < 2 {
	// 	fmt.Println("Please enter the queue name that you want to subscribe!")
	// 	return
	// }
	subQueName := "FIRST_QUEUE"

	var name string
	var timeout int64
	flag.StringVar(&name, "n", "", subQueName)
	flag.Int64Var(&timeout, "t", 10, "(Optional) Timeout in seconds for long polling")
	flag.Parse()

	if len(name) == 0 {
		flag.PrintDefaults()
	}

	fmt.Println("Start..")

	//Create a session
	conf := aws.Config{Region: aws.String("us-west-2")}
	sess := session.New(&conf)

	//Create a SQS service client
	svc := sqs.New(sess)

	resultURL, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(subQueName),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == sqs.ErrCodeQueueDoesNotExist {
			fmt.Println("Unable to find queue %q.", name)
		}
		fmt.Println("Unable to queue %q, %v.", name, err)
	}

	//Read message
	for {

		result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: resultURL.QueueUrl,
			AttributeNames: aws.StringSlice([]string{
				"SentTimestamp",
			}),
			MaxNumberOfMessages: aws.Int64(1),
			MessageAttributeNames: aws.StringSlice([]string{
				"All",
			}),
			WaitTimeSeconds: aws.Int64(timeout),
		})

		if err != nil {
			fmt.Println("Unable to receive message from queue %q, %v.", name, err)
		}

		fmt.Printf("Received %d messages.\n", len(result.Messages))
		if len(result.Messages) > 0 {
			fmt.Println(result.Messages)
		}
	}
}
