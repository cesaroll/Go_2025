package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Actor struct {
	Id        int64
	FirstName string
	LastName  string
}

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

	// Add an newActor
	var newActor = Actor{
		FirstName: "Tom",
		LastName:  "Hanks",
	}
	ctx := context.Background()
	actorId, err := addActor(ctx, newActor)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Id of added actor: %v\n", actorId)

	// Read an actor
	actor, err := getActor(actorId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor: %v\n", actor)

	// Re-add actor would fail (to test transaction)
	ctx = context.Background()
	_, err = addActor(ctx, newActor)
	if err != nil {
		log.Fatal(err)
	}

	// Update an actor
	actor.FirstName = "Tommy"
	err = updateActor(actor)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor name updated: %v\n", actor)

	// Read all actors
	actors, err := getAllActors()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actors: %v\n", actors)

	// Delete an actor
	err = deleteActor(actor.Id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Actor deleted %v\n", actor)
}

func addActor(ctx context.Context, actor Actor) (int64, error) {

	// Verify if actor exists before adding
	// Begin transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	defer tx.Rollback()

	// Check if name exists
	var id int64
	err = tx.QueryRowContext(ctx, "SELECT id FROM actors WHERE first_name = ? AND last_name = ?",
		actor.FirstName, actor.LastName).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Actor does not exists")
		} else {
			return 0, fmt.Errorf("addActor: %v", err)
		}
	}

	// Rollback if actor exists
	if id > 0 {
		if err = tx.Rollback(); err != nil {
			return 0, fmt.Errorf("addActor: %v", err)
		}
		fmt.Println("Actor already exists", id)
		fmt.Println("*** Transaction rolled back ***")
		return id, nil
	}

	// Add actor
	result, err := db.Exec("INSERT INTO actors (first_name, last_name) VALUES (?, ?)",
		actor.FirstName, actor.LastName)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	return id, nil
}

func getActor(id int64) (Actor, error) {
	var firstName, lastName string
	var actor Actor
	err := db.QueryRow("SELECT first_name, last_name FROM actors WHERE id = ?", id).Scan(&firstName, &lastName)
	if err != nil {
		return actor, fmt.Errorf("readActor: %v", err)
	}
	return Actor{Id: id, FirstName: firstName, LastName: lastName}, nil
}

func getAllActors() ([]Actor, error) {
	var actors []Actor
	rows, err := db.Query("SELECT id, first_name, last_name FROM actors")
	if err != nil {
		return nil, fmt.Errorf("getAllActors: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var actor Actor
		err := rows.Scan(&actor.Id, &actor.FirstName, &actor.LastName)
		if err != nil {
			return nil, fmt.Errorf("getAllActors: %v", err)
		}
		actors = append(actors, actor)
	}
	return actors, nil
}

func updateActor(actor Actor) error {
	_, err := db.Exec("UPDATE actors SET first_name = ?, last_name = ? WHERE id = ?",
		actor.FirstName, actor.LastName, actor.Id)
	if err != nil {
		return fmt.Errorf("updateActor: %v", err)
	}
	return nil
}

func deleteActor(id int64) error {
	_, err := db.Exec("DELETE FROM actors WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("deleteActor: %v", err)
	}
	return nil
}
