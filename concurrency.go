package main

import (
	"fmt"
	"time"
)

// Define a function that simulates work by sleeping for a given duration, then sends a message on a channel when finished
func worker(id int, workDuration time.Duration, c chan int) {
	fmt.Printf("Worker %d: Starting work for %v seconds\n", id, workDuration.Seconds())
	time.Sleep(workDuration)
	c <- id
}

func main11() {
	// Create a buffered channel with capacity equal to the number of workers.
	// This means the channel can hold a value for each worker without needing a receiver ready.
	c := make(chan int, 5)

	// Launch several workers as goroutines, each with a different work duration.
	// Each worker will send on the channel when finished.
	for i := 1; i <= 5; i++ {
		go worker(i, time.Duration(i)*time.Second, c)
	}

	// Close the channel after all values have been sent.
	// This must be done in a separate goroutine, because the sends are not yet complete (they're sleeping).
	go func() {
		// This waits for all sends on the channel to complete, then closes the channel.
		// It's safe to do this because the senders (the worker goroutines) have no references to the channel after they finish sending.
		for i := 1; i <= 5; i++ {
			<-c
		}
		close(c)
	}()

	// Receive from the channel until it's closed and empty.
	// This waits for all workers to finish.
	// If the workers finish in a different order than they were launched, that's the order we receive the values in.
	for id := range c {
		fmt.Printf("Worker %d: Finished\n", id)
	}
}
