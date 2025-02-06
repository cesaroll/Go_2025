package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for range time.Tick(500 * time.Millisecond) {
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			println("tick!")
		}
	}(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()

}
