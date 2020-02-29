package followers

import(
	"log"
	"encoding/json"
)

func Followers(userId int) []int{

	db := Database()
	result,err := db.Query("select * from fitnessusers where id = $1",userId)
			if err != nil {
				log.Fatal(err)
			}
			var m []int
			usr := new(User)
			for result.Next(){
				
				err:=result.Scan(&usr.ID,&usr.NAME,&usr.FOLLOWERS,&usr.FOLLOWING)
				if err!=nil{
					panic(err.Error())
				}
				
				log.Println("program ran successfully",result)
			}
			json.Unmarshal(usr.FOLLOWERS,&m)
			log.Println("Followers",m)
			return m
		}