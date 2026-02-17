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

```
