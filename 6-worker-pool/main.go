package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worked %d started job %d\n", id, job)
		results <- job * 2                // simulate work by multipling by 2
		time.Sleep(20 * time.Millisecond) // ensure different workers for contiguos jobs
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup

	// start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// send jobs to workers
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // no more jobs to send

	wg.Wait()
	close(results)

	// collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
