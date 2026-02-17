package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mutex   sync.Mutex // declare mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	mutex.Lock() // acquire lock before accessing shared variable
	counter++
	mutex.Unlock() // release lock
}

func main() {
	start := time.Now()
	numGoroutines := 100
	wg.Add(numGoroutines)

	for range numGoroutines {
		go increment()
	}
	wg.Wait()

	fmt.Println("Final counter:", counter)
	fmt.Println("Elapsed time:", time.Since(start))
}
