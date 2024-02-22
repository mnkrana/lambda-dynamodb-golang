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
	KEY_UUID         = "uuid"
	KEY_ConnectionID = "connectionID"
	KEY_State        = "clientState"
)

type ConnectionItem struct {
	UUID         string `json:"uuid"`
	ConnectionID string `json:"connectionID"`
	State        int    `json:"clientState"`
}
