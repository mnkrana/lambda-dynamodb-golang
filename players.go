package lambda_dynamodb_golang

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type PlayerInfo struct {
	UUID    string `json:"uuid"`
	device  string `json:"device"`
	session int    `json:"session"`
	online  int    `json:"online"`
}

func PutNewPlayer(uuid string) {
	playerInfo := PlayerInfo{
		UUID:    uuid,
		device:  "Empty",
		session: 1,
		online:  0,
	}

	attributeValues, _ := dynamodbattribute.MarshalMap(playerInfo)

	input := &dynamodb.PutItemInput{
		Item:      attributeValues,
		TableName: aws.String(TABLE),
	}

	_, err := dynamodbSession.PutItem(input)
	if err != nil {
		log.Printf("Error in puting item %v", err)
	}
}
