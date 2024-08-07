package lambda_dynamodb_golang

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type PlayerInfo struct {
	UUID         string `json:"uuid"`
	PlayerDevice string `json:"player_device"`
	SessionCount int    `json:"session_count"`
}

func PutNewPlayer(uuid string, playerDevice string) {
	playerInfo := PlayerInfo{
		UUID:         uuid,
		PlayerDevice: playerDevice,
		SessionCount: 1,
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

func GetPlayerInfo(uuid string) PlayerInfo {
	return findPlayerByKeyValue("uuid", uuid)
}
