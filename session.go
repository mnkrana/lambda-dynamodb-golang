package lambda_dynamodb_golang

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamodbSession *dynamodb.DynamoDB

func InitSession() {
	if dynamodbSession == nil {
		dynamodbSession = newDynamoDBSession()
	}
}

func newDynamoDBSession() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})
	return dynamodb.New(sess)
}
