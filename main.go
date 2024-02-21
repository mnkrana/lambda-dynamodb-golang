package lambda_dynamodb_golang

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

func Init() {
	getAWSConfig()
	initSession()
}

func PutNewItem(item string) {

	connectionItem := ConnectionItem{
		UUID:         uuid.New().String(),
		ConnectionID: item,
	}

	attributeValues, _ := dynamodbattribute.MarshalMap(connectionItem)

	input := &dynamodb.PutItemInput{
		Item:      attributeValues,
		TableName: aws.String(TABLE),
	}

	_, err := dynamodbSession.PutItem(input)
	if err != nil {
		log.Printf("Error in puting item %v", err)
	}
}

func DeleteItemByKeyValue(key string, value string) {
	filt := expression.Name(key).Equal(expression.Value(value))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(TABLE),
	}

	result, err := dynamodbSession.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}

	//NOTE: delete only first found item at 0 index
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(*result.Items[0]["uuid"].S),
			},
		},
		TableName: aws.String(TABLE),
	}

	_, err = dynamodbSession.DeleteItem(input)

	if err != nil {
		log.Printf("Error in deleting item %v", err)
	} else {
		log.Printf("Successfully deleted item")
	}
}
