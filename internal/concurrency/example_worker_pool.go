package concurrency

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // stops when jobs closed
		fmt.Println("worker", id, "processing", job)
		results <- job * 2
	}
}

func WorkerPoolMainExample() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// start workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // signal no more jobs

	// collect results
	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}
}
