package main

import (
    "go.mongodb.org/mongo-driver/v2/mongo"
)

type actor struct {
	FirstName	string
	LastName	string
	Awards		int16
}

func getActorsCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("dvdstore").Collection("actors")
	return collection
}