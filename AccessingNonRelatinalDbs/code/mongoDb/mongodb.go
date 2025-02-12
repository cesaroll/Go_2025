package mongoDb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongo(uri string) (*mongo.Client, error) {

	// Set client Options
	clientOptions := options.Client().ApplyURI(uri).
		SetAuth(options.Credential{
			Username: "admin",
			Password: "password",
		})

	// Connect to mongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println(">> Connected to MongoDB!")
	return client, nil
}

func DisconnectMongo(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("** Disconnected from MongoDB!")
	}
}
