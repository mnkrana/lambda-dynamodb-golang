package lambda_dynamodb_golang

import (
	"log"
	"strconv"

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

func PutNewItem(connectionID string) {

	connectionItem := ConnectionItem{
		UUID:         uuid.New().String(),
		ConnectionID: connectionID,
		State:        int(Ready),
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

func UpdateItem(key string, value string, connectionID string, state State) {

	primaryKey := findPrimaryKeyValue(key, value)

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(primaryKey),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newState": {
				N: aws.String(strconv.Itoa(state.EnumIndex())),
			},
		},
		UpdateExpression: aws.String("SET state = :newState"),
		TableName:        aws.String(TABLE),
	}

	_, err := dynamodbSession.UpdateItem(input)
	if err != nil {
		log.Printf("Error in updaing item %v", err)
	}
}

func DeleteItemByKeyValue(key string, value string) {
	primaryKey := findPrimaryKeyValue(key, value)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(primaryKey),
			},
		},
		TableName: aws.String(TABLE),
	}

	_, err := dynamodbSession.DeleteItem(input)

	if err != nil {
		log.Printf("Error in deleting item %v", err)
	} else {
		log.Printf("Successfully deleted item")
	}
}

func findPrimaryKeyValue(key string, value string) string {
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

	return *result.Items[0]["uuid"].S
}
