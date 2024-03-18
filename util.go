package lambda_dynamodb_golang

import (
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func findItemByKeyValue(key string, value string) ConnectionItem {
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
	state, err := strconv.Atoi(*result.Items[0][KEY_State].N)
	if err != nil {
		log.Println("Error in converting state!")
	}

	player, err := strconv.Atoi(*result.Items[0][KEY_Player].N)
	if err != nil {
		log.Println("Error in converting player!")
	}
	item := ConnectionItem{
		UUID:               *result.Items[0][KEY_UUID].S,
		MyConnectionID:     *result.Items[0][KEY_MyConnectionID].S,
		FriendConnectionID: *result.Items[0][KEY_FriendConnectionID].S,
		State:              state,
		Player:             player,
	}
	return item
}

func findPlayerByKeyValue(key string, value string) PlayerInfo {
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
		return PlayerInfo{}
	}

	//return first find
	return getPlayerFromResult(result)
}

func getPlayerFromResult(result *dynamodb.ScanOutput) PlayerInfo {
	session, err := strconv.Atoi(*result.Items[0]["session_count"].N)
	if err != nil {
		log.Println("Error in converting state!")
	}

	online, err := strconv.Atoi(*result.Items[0]["is_online"].N)
	if err != nil {
		log.Println("Error in converting player!")
	}
	playerInfo := PlayerInfo{
		UUID:         *result.Items[0]["uuid"].S,
		PlayerDevice: *result.Items[0]["player_device"].S,
		SessionCount: session,
		IsOnline:     online,
	}
	return playerInfo
}
