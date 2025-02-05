package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	receivedOrdersCh := receiveOrders()
	validOrdersCh, invalidOrdersCh := validateOrders(receivedOrdersCh)

	wg.Add(1)
	go func(validOrdersCh <-chan order, invalidOrdersCh <-chan invalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validOrdersCh:
				if ok {
					fmt.Printf("Valid order received: %v\n", order)
				} else {
					break loop
				}
			case invalidOrder, ok := <-invalidOrdersCh:
				if ok {
					fmt.Printf("Invalid order received: %v. Issue: %v\n", invalidOrder.Order, invalidOrder.Error)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validOrdersCh, invalidOrdersCh)

	wg.Wait()

}

func validateOrders(in <-chan order) (<-chan order, <-chan invalidOrder) {
	valid := make(chan order)
	invalid := make(chan invalidOrder, 1)

	go func(valid chan<- order, invalid chan<- invalidOrder) {
		for order := range in {
			if order.Quantity <= 0 {
				invalid <- invalidOrder{Order: order, Error: errors.New("quantity must be greater than 0")}
			} else {
				valid <- order
			}
		}
		close(valid)
		close(invalid)
	}(valid, invalid)

	return valid, invalid
}

func receiveOrders() <-chan order {
	out := make(chan order)

	go func(out chan<- order) {
		for _, rawOrder := range rawOrders {
			var newOrder order
			err := json.Unmarshal([]byte(rawOrder), &newOrder)
			if err != nil {
				log.Print(err)
				continue
			}
			out <- newOrder
			// fmt.Printf("New Order Received: %v\n", newOrder)
		}
		close(out)
	}(out)

	return out
}

var rawOrders = []string{
	`{"ProductCode": 1111, "Quantity": -5, "Status": 1}`,
	`{"ProductCode": 2222, "Quantity": 10, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 15, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 20, "Status": 1}`,
}
