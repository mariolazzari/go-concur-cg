# Go concurrency by the Coding Gopher

[YouTube playlist](https://www.youtube.com/watch?v=3MY0B5PBgR8&list=PLqR6Wq9GKBiunQOAao0_Sjda0Kvp5TqGO&index=9)

## Goroutines

 Lightweight, efficient threads managed by the Go runtime, allowing non-blocking, parallel tasks.

```go
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := range 5 {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	start := time.Now()

	go printMessage("Goroutine 1")
	go printMessage("Goroutine 2")

	time.Sleep(2100 * time.Millisecond)
	fmt.Println("Elapsed time:", time.Since(start))
}
```

## Channels (bidirectional)

Can both send and receive data

```go
package main

import (
	"fmt"
	"time"
)

func senderReceiver(ch chan string) {
	fmt.Println("Sending message...")
	ch <- "Message from sender" // Send message to main

	response := <-ch // Receive message from main
	fmt.Println("Received in Goroutine:", response)
}

func main() {
	start := time.Now()

	ch := make(chan string)

	// start goroutine
	go senderReceiver(ch)

	// receive message from sender
	msg := <-ch
	fmt.Println("Received in Main:", msg)

	// send a response back to goroutine
	time.Sleep(time.Second)
	ch <- "Response from main"
	time.Sleep(time.Second)

	fmt.Println("Elapsed time:", time.Since(start))
}
```

## Unidirectional Channels

Either send or receive data but not both, making your concurrent programs more efficient and clear.

```go
package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- string) {
	fmt.Println("Sending message...")
	ch <- "Message from sender" // Send message to main
}

func receiver(ch <-chan string) {
	response := <-ch // Receive message from main
	fmt.Println("Received in Goroutine:", response)
}

func main() {
	start := time.Now()

	ch := make(chan string)

	// start goroutine
	go sender(ch)
	go receiver(ch)

	time.Sleep(time.Second)

	fmt.Println("Elapsed time:", time.Since(start))
}
```

## Buffered channels

Buffered channels allow Goroutines to send multiple messages without blocking, improving the efficiency of your concurrent programs.

```go
package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "Message 1"
	ch <- "Message 2"
	ch <- "Message 3"
	close(ch) // close channel to avoid deadlock
}

func main() {
	start := time.Now()

	ch := make(chan string, 3) // buffered channel
	go sendData(ch)

	// receive data from channel
	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("Elapsed time:", time.Since(start))
}
```

## Mutex

Use mutexes in Go for synchronizing goroutines and preventing race conditions in concurrent programming.

```go
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
	mutex.Lock()         // acquire lock before accessing shared variable
	defer mutex.Unlock() // defer unlock execution in order to release lock on error too

	counter++
	// mutex.Unlock() // release lock
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
```

## Worker pool pattern

Worker pools are a powerful way to limit the number of goroutines processing tasks simultaneously, improving performance and resource efficiency in your Go programs. 

```go
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
```

## Concurrent url fetcher
