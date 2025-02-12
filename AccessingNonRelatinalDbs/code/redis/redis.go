package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "k16SKsVuGx",
		DB:       0,
	})
	defer client.Close()

	pong, err := client.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	fmt.Println("Connected to Redis")

}
