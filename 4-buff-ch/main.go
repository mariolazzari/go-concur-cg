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
