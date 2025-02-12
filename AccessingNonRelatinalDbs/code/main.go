package main

import (
	"context"
	"fmt"
	"log"

	"nosql/mongoDb"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

	actorsCollection := getActorsCollection(client)

	// Insert One
	insertOne(actorsCollection, actor{"Tom", "Cruise", 7})

	// Insert Many
	mark := actor{"Mark", "Hamil", 2}
	mili := actor{"Mili", "Bobby Brown", 3}
	insertMany(actorsCollection, mark, mili)

	// Retrieve Actor
	retrieveOne(actorsCollection, "Mili", "Bobby Brown")

	// Update One Actor
	updateActorAwards(actorsCollection, "Tom", "Cruise", 8)

	// Retrieve Many
	retrieveMany(actorsCollection, "Cruise")

}

func updateActorAwards(collection *mongo.Collection, firstName string, lastName string, awards int16) {

	filter := bson.D{{"firstname", firstName}, {"lastname", lastName}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "awards", Value: awards},
		}},
	}

	updateResult, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v actors and updated %v actors\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func retrieveMany(collection *mongo.Collection, lastName string) []actor {
	var results []actor

	filter := bson.D{{"lastname", lastName}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result actor
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	fmt.Println("Actors Retrieved: ", results)
	return results
}

func retrieveOne(collection *mongo.Collection, firstName string, lastName string) actor {
	var result actor

	// filter := bson.D{}
	filter := bson.D{{"firstname", firstName}, {"lastname", lastName}}

	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Actor Retrieved: ", result)
	return result
}

func insertOne(collection *mongo.Collection, actor actor) {
	insertResult, err := collection.InsertOne(context.TODO(), actor)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted new Actor: ", insertResult.InsertedID)
}

func insertMany(collection *mongo.Collection, actors ...actor) {
	insertManyResult, err := collection.InsertMany(context.TODO(), actors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple Actors: ", insertManyResult.InsertedIDs)
}
