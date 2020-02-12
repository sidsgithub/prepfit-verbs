package mainMods

import (
	"log"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func Database(prog Program,s string){
	log.Println(s)
	db, err := sql.Open("mysql","root:password@tcp(localhost:3306)/customerOne")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s)
	switch s{
		case "findProgram":
			result,err := db.Query("select * from program where id = ?",prog.ID)
			if err != nil {
				log.Fatal(err)
			}
			for result.Next(){
				err:=result.Scan(&prog.ID,&prog.TITLE,&prog.DESC)
				if err!=nil{
					panic(err.Error())
				}
				log.Println("program ran successfully",prog.TITLE)
			}
		case "createProgram":
			row, err := db.Query("INSERT INTO program(id,title,des) values (?,?,?);",prog.ID,prog.TITLE,prog.DESC)
			if err != nil {
						log.Fatal(err)
					}
					log.Println(row,"inserted")
		case "dummystring":
			log.Println("")

	}
	defer db.Close()
}