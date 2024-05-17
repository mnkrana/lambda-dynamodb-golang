package lambda_dynamodb_golang

type Player int

const (
	Me Player = iota + 1
	Friend
)

func (p Player) EnumIndex() int {
	return int(p)
}

const (
	KEY_UUID               = "uuid"
	KEY_MyConnectionID     = "myConnectionID"
	KEY_FriendConnectionID = "friendConnectionID"
	KEY_State              = "clientState"
	KEY_Player             = "player"
	KEY_PlayerUUID         = "playerUUID"
)

type ConnectionItem struct {
	UUID               string `json:"uuid"`
	MyConnectionID     string `json:"myConnectionID"`
	FriendConnectionID string `json:"friendConnectionID"`
	State              int    `json:"clientState"`
	Player             int    `json:"player"`
	PlayerUUID         string `json:"playerUUID"`
}
