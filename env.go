package lambda_dynamodb_golang

import "os"

const (
	KEY_REGION = "dynamodb-region"
	KEY_TABLE  = "dynamodb-table"
)

var (
	REGION string
	TABLE  string
)

func GetAWSConfig() {
	REGION = os.Getenv(KEY_REGION)
	TABLE = os.Getenv(KEY_TABLE)
}
