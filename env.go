package lambda_dynamodb_golang

import "os"

const (
	KEY_REGION            = "dynamodb_region"
	KEY_TABLE             = "dynamodb_table"
	KEY_GAME_CONFIG_TABLE = "dynamodb_game_config_table"
)

var (
	REGION            string
	TABLE             string
	GAME_CONFIG_TABLE string
)

func getAWSConfig() {
	if len(REGION) == 0 {
		REGION = os.Getenv(KEY_REGION)
		TABLE = os.Getenv(KEY_TABLE)
		GAME_CONFIG_TABLE = os.Getenv(KEY_GAME_CONFIG_TABLE)
	}
}
