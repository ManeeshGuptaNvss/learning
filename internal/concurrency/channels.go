package concurrency

import "fmt"

/*
Reference : https://youtu.be/5Z8skvm4g64?si=XAYq7x51f2C_DRr6
*/

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
