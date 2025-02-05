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
	reservedInventoryCh := reserveInventory(validOrdersCh)
	fillOrders(reservedInventoryCh, &wg)

	wg.Add(1)

	go func(invalidOrdersCh <-chan invalidOrder) {
		for order := range invalidOrdersCh {
			fmt.Printf("Invalid order received: %v, Issue: %v\n", order.Order, order.Error)
		}
		wg.Done()
	}(invalidOrdersCh)

	wg.Wait()

}

func fillOrders(in <-chan order, wg *sync.WaitGroup) {
	const workers = 3
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			for order := range in {
				order.Status = filled
				fmt.Printf("Order has been completed: %v\n", order)
			}
			wg.Done()
		}()
	}
}

func reserveInventory(in <-chan order) <-chan order {
	out := make(chan order)
	var wg sync.WaitGroup

	const workers = 3
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for order := range in {
				order.Status = reserved
				out <- order
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
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
