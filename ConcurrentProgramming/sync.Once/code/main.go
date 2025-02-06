package main

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	Run()
	Run()
}

var db *sql.DB
var o sync.Once

func Run() {
	o.Do(func() {
		log.Println("Opening database connection")
		var err error
		db, err = sql.Open("sqlite3", "./mydb.db")
		if err != nil {
			log.Fatal(err)
		}
	})
}
