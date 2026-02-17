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
