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

func PutNewPlayer(uuid string, device string) {
	log.Printf("add new player with %v in db", uuid)
	playerInfo := PlayerInfo{
		UUID:    uuid,
		device:  device,
		session: 1,
		online:  1,
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

func UpdatePlayerByUUIDN(uuid string, key string, value string, action string) {
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
