package lambda_dynamodb_golang

import "os"

const (
	KEY_REGION = "dynamodb_region"
	KEY_TABLE  = "dynamodb_table"
)

var (
	REGION string
	TABLE  string
)

func getAWSConfig() {
	REGION = os.Getenv(KEY_REGION)
	TABLE = os.Getenv(KEY_TABLE)
}
