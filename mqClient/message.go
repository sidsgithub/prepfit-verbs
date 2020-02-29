package mqClient

type Message struct {
	USERID string `json:"userId" description:"identifier"`
	ACTION string `json:"action" description:"action "`
} 