package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	// Set client Options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:32017")

	// Connect to mongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(">> Connected to MongoDB!")

	// Close the connection
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("** Disconnected from MongoDB!")
	}

}
