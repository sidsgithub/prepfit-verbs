package followers

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
)

const (
  host     = "127.0.0.1"
  port     = 5432
  user     = "sid"
  password = "password"
  dbname   = "testdb"
)

func Database()(*sql.DB){

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	
  db, err := sql.Open("postgres", psqlInfo)

  if err != nil {
    panic(err)
  }
  err = db.Ping()

  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  return db
}