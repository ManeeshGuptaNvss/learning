package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ConcurrentPrint() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {

		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Printf("GO ROUTINE 1: IDX %v \n", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Printf("GO ROUTINE 2: IDX %v \n", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	// this will stops the main routine until the above 2 go routines finishes
	wg.Wait()
	fmt.Println("DONE!!!")
}
