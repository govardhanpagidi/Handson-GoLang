package DB

import (
	"fmt"
	
)

func SaveEvent(string key, string value){

}

func ConnectDynamoDB(string conString){

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("Mumbai")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	return true;
}