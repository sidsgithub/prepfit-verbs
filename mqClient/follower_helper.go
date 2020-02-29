package mqClient

import (
	"log"
	"strconv"
	"github.com/user/prepFitness/followers"
)

func Followers_helper(user_id int, followers_ []int) {

	id := strconv.Itoa(user_id)

	for i := 0; i < len(followers_); i++ {
		msg := `user : ` + id + ` updated a program`
		// notification := Notification{followers_[i], msg}
		db := followers.Database()
		
		row, err := db.Query("insert into notifications(t_user,t_msg) values($1,$2);",followers_[i],msg)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(row, "inserted")
		// go SendHelper(notification)
	}
}
