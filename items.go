package lambda_dynamodb_golang

type State int

const (
	Ready State = iota + 1
	Playing
)

func (d State) EnumIndex() int {
	return int(d)
}

type ConnectionItem struct {
	UUID         string `json:"uuid"`
	ConnectionID string `json:"connectionID"`
	State        int    `json:"state"`
}
