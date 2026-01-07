package concurrency

import (
	"context"
	"fmt"
	"time"
)

/*
Producer
- Generate an incremental number every 700 milliseconds.
- Exit on its own after 4 seconds
Consumer
- Print generated number
- Exit after printing all numbers

*/

func RunProducerConsumer() {

	numbers := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)

	defer cancel()

	go func(ctx context.Context, ch chan<- int) {
		i := 0
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- i:
				i++
				time.Sleep(700 * time.Millisecond)

			}
		}
	}(ctx, numbers)

	func(ch <-chan int) {
		for num := range ch {
			fmt.Println(num)
		}
	}(numbers)
}
