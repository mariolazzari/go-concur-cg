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
