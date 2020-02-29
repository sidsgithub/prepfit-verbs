package followers
import (
	// "encoding/json"
)

type User struct {
	ID        int    `json:"id" description:"identifier"`
	NAME      string `json:"name" description:"identifier"`
	FOLLOWERS []byte `json:"followers" description:"identifier"`
	FOLLOWING []byte `json:"following" description:"identifier"`
}
