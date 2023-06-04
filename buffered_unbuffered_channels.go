package main

import (
	"fmt"
	"time"
)

func main14() {
	// Unbuffered channel
	unbuffered := make(chan string)
	go func() {
		unbuffered <- "hello"
		fmt.Println("Sent to unbuffered channel")
	}()
	time.Sleep(time.Second) // Wait for the sender goroutine to run
	fmt.Println("Received from unbuffered channel:", <-unbuffered)

	// Buffered channel
	buffered := make(chan string, 1) // Create a buffered channel with capacity 1
	go func() {
		buffered <- "world"
		fmt.Println("Sent to buffered channel")
	}()
	fmt.Println("Received from buffered channel:", <-buffered)
}
