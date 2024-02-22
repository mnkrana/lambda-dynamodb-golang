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
		UUID:               uuid.New().String(),
		MyConnectionID:     connectionID,
		FriendConnectionID: "",
		State:              int(Ready),
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

func UpdateStateByKeyValue(key string, value string, connectionID string, state State) {

	item := findItemByKeyValue(key, value)

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			KEY_UUID: {
				S: aws.String(item.UUID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newState": {
				N: aws.String(strconv.Itoa(state.EnumIndex())),
			},
		},
		UpdateExpression: aws.String("SET " + KEY_State + " = :newState"),
		TableName:        aws.String(TABLE),
	}

	_, err := dynamodbSession.UpdateItem(input)
	if err != nil {
		log.Printf("Error in updaing item %v", err)
	}
}

func UpdateItemByUUID(uuid string, key string, value string, action string) {
	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			KEY_UUID: {
				S: aws.String(uuid),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newState": {
				N: aws.String(value),
			},
		},
		UpdateExpression: aws.String(action + " " + key + " = :newState"),
		TableName:        aws.String(TABLE),
	}

	_, err := dynamodbSession.UpdateItem(input)
	if err != nil {
		log.Printf("Error in updaing item %v", err)
	}
}

func DeleteItemByKeyValue(key string, value string) {
	item := findItemByKeyValue(key, value)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			KEY_UUID: {
				S: aws.String(item.UUID),
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

// return finder's uuid and other found uuid
func FindOtherReadyItem(key string, value string) (ConnectionItem, ConnectionItem, bool) {
	item := findItemByKeyValue(key, value)

	filt1 := expression.Name(KEY_UUID).NotEqual(expression.Value(item.UUID))
	filt2 := expression.Name(KEY_State).Equal(expression.Value(Ready.EnumIndex()))

	expr, err := expression.NewBuilder().WithFilter(filt1.And(filt2)).Build()

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
		log.Printf("Query API call failed: %s", err)
	}

	if *result.Count == 0 {
		log.Println("No ready item found")
		return item, ConnectionItem{}, false
	}

	//return first find
	return item, getConnectionItemFromResult(result), true
}
