package lambda_dynamodb_golang

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type PlayerInfo struct {
	UUID          string `json:"uuid"`
	player_device string `json:"player_device"`
	session_count int    `json:"session_count"`
	is_online     int    `json:"is_online"`
}

func PutNewPlayer(uuid string, playerDevice string) {
	playerInfo := PlayerInfo{
		UUID:          uuid,
		player_device: playerDevice,
		session_count: 1,
		is_online:     1,
	}

	attributeValues, err := dynamodbattribute.MarshalMap(playerInfo)

	if err != nil {
		log.Printf("Error in marshal map %v output is %v", playerInfo, attributeValues)
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValues,
		TableName: aws.String(TABLE),
	}

	_, err = dynamodbSession.PutItem(input)
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
