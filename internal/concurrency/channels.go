package concurrency

import "fmt"

func PrintSumUsingChannels() {
	c := make(chan int)

	go func() {
		sum := 0
		for i := 0; i <= 100; i++ {
			sum += i
		}
		c <- sum
	}()
	output := <-c
	fmt.Printf("PrintSumUsingChannels(): %v \n", output)

}
