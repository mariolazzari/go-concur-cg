package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Result struct {
	URL      string
	Duration time.Duration
	Error    error
}

func fetchURL(url string, results chan<- Result) {
	start := time.Now()

	// random delay to simulate network latency
	delay := time.Duration(rand.IntN(3000)+1000) * time.Millisecond
	time.Sleep(delay)

	duration := time.Since(start)

	// randonm error (30%)
	var err error
	if rand.IntN(10) < 3 {
		err = fmt.Errorf("failed to fetch %s", url)
	}

	results <- Result{URL: url, Duration: duration, Error: err}
}

func main() {
	urls := []string{
		"https://mariolazzari.it",
		"https://golang.org",
		"https://google.com",
	}

	var wg sync.WaitGroup
	results := make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			fetchURL(url, results)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Fetching urls cocnurrently...")

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error fetching %s: %v\n", result.URL, result.Error)
		} else {
			fmt.Printf("Fetched %s: %v\n", result.URL, result.Duration)
		}
	}
}
