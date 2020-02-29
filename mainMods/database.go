package mainMods

import (
	"log"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
func Database()(*sql.DB){
	db, err := sql.Open("mysql","root:password@tcp(localhost:3306)/customerOne")
	if err != nil {
		log.Fatal(err)
	}
	return db
}