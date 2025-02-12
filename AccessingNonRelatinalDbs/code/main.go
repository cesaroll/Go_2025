package main

import (
	"context"
	"fmt"
	"log"

	"nosql/mongoDb"
)

func main() {

	// Define MongoDB URI
	mongoURI := "mongodb://localhost:32017"

	// Connect to MongoDB
	client, err := mongoDb.ConnectMongo(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoDb.DisconnectMongo(client) // Ensures disconnection when the program exits

	// Insert One
	tom := actor{"Tom", "Hanks", 9}
	actorsCollection := getActorsCollection(client)
	insertResult, err := actorsCollection.InsertOne(context.TODO(), tom)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted new Actor: ", insertResult.InsertedID)

	// Insert Many
	mark := actor{"Mark", "Hamil", 2}
	mili := actor{"Mili", "Bobby Brown", 3}
	actors := []interface{}{mark, mili}
	insertManyResult, err := actorsCollection.InsertMany(context.TODO(), actors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple Actors: ", insertManyResult.InsertedIDs)

}
