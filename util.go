package lambda_dynamodb_golang

import (
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func findItemByKeyValue(key string, value string) ConnectionItem {
	log.Printf("Find Primary key value of %v", value)

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
		log.Printf("Query API call failed: %s", err)
	}

	if *result.Count == 0 {
		log.Println("No key found")
		return ConnectionItem{}
	}

	//return first find
	return getConnectionItemFromResult(result)
}

func getConnectionItemFromResult(result *dynamodb.ScanOutput) ConnectionItem {

	log.Printf("UUID %v", *result.Items[0][KEY_UUID].S)
	log.Printf("MyConnectionID %v", *result.Items[0][KEY_MyConnectionID].S)
	log.Printf("FriendConnectionID %v", *result.Items[0][KEY_FriendConnectionID].S)

	state, err := strconv.Atoi(*result.Items[0][KEY_State].S)
	if err != nil {
		log.Println("Error in converting state!")
	}
	log.Printf("State %v", state)

	item := ConnectionItem{
		UUID:               *result.Items[0][KEY_UUID].S,
		MyConnectionID:     *result.Items[0][KEY_MyConnectionID].S,
		FriendConnectionID: *result.Items[0][KEY_FriendConnectionID].S,
		State:              state,
	}
	return item
}