package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	receiveOrders()
	fmt.Println(orders)
}

func receiveOrders() {

	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		orders = append(orders, newOrder)
	}
}

var rawOrders = []string{
	`{"ProductCode": 1111, "Quantity": 5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 10, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 15, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 20, "Status": 1}`,
}
