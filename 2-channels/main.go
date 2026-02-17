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
