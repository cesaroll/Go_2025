package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Data Source name properties
	dsn := mysql.Config{
		User:   "root",
		Passwd: "start123",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "movies",
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	// Add an actor
	actorId, err := addActor("Tom", "Hanks")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id of added actor: %v\n", actorId)
}

func addActor(firstName string, lastName string) (int64, error) {
	result, err := db.Exec("INSERT INTO actors (first_name, last_name) VALUES (?, ?)", firstName, lastName)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	return id, nil
}
