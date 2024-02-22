package lambda_dynamodb_golang

type State int

const (
	Ready State = iota + 1
	Playing
)

func (d State) EnumIndex() int {
	return int(d)
}

const (
	KEY_UUID               = "uuid"
	KEY_MyConnectionID     = "myConnectionID"
	KEY_FriendConnectionID = "friendConnectionID"
	KEY_State              = "clientState"
)

type ConnectionItem struct {
	UUID               string `json:"uuid"`
	MyConnectionID     string `json:"myConnectionID"`
	FriendConnectionID string `json:"friendConnectionID"`
	State              int    `json:"clientState"`
}
