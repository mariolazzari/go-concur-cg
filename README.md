# Go concurrency by the Coding Gopher

[YouTube](https://www.youtube.com/watch?v=3MY0B5PBgR8&list=PLqR6Wq9GKBiunQOAao0_Sjda0Kvp5TqGO&index=9)

## Goroutines

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

## Unidirectional Channels

```go
package main

import (
	"fmt"
	"time"
)

// unidirectional channel for sending data
func sendData(sendOnly chan<- string) {
	sendOnly <- "hello from sendData"
}

// unidirectional channel for receiving data
func receiveData(receiveOnly <-chan string) {
	msg := <-receiveOnly
	fmt.Println("Received:", msg)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	// start goroutine sending data
	go sendData(ch)

	// start goroutine receiving data
	go receiveData(ch)

	time.Sleep(time.Second)
	fmt.Println("Elapsed time:", time.Since(start))
}
```

## Buffered channels

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
```
