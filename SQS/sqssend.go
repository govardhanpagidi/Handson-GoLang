package main
import(
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws/credentials"
)


// Source : https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sqs-example-create-queue.html

// Prerequisites 
//Create a IAM user in AWS and use credentials to create  a session

//How to run
func main(){

	if len(os.Args) < 2 {
		fmt.Println("Please enter some messge to send to queue..")
		return
	}

	message := os.Args[1]
	fmt.Println("Start!")



	//Create AWS session
	conf := aws.Config{Region: aws.String("us-west-2")}
	sess := session.New(&conf)

	
	//Create a SQS service client
	svc:= sqs.New(sess)


	qURL := "https://sqs.us-west-2.amazonaws.com/771658639958/FIRST_QUEUE"

	//Prepare and send message to queue
	result,err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds : aws.Int64(10),
		MessageAttributes : map[string]*sqs.MessageAttributeValue{
			"Title" : &sqs.MessageAttributeValue{
				DataType: aws.String("String"),
				StringValue : aws.String("Test-Title"),
			},
		},
		MessageBody : aws.String(message),
		QueueUrl : &qURL,
	})
	
	if err != nil {
		fmt.Println("Error : ",err)
	}

	fmt.Println("Success", *result.MessageId)

	fmt.Println("End")
}