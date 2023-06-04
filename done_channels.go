package main

import (
	"fmt"
	"sync"
)

// A function that simulates work, signaling completion on a done channel.
func worker2(id int, done chan bool) {
	fmt.Printf("Worker %d: Starting work...\n", id)
	// Simulating work with a sleep. In real-world usage this would be replaced by actual work.
	for i := 0; i <= id; i++ {
	}
	fmt.Printf("Worker %d: Finished work\n", id)
	done <- true
}

func main15() {
	numWorkers := 5
	done := make(chan bool) // Channel to signal when a worker has finished.
	var wg sync.WaitGroup   // WaitGroup to wait for all workers to finish.

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			worker2(id, done)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	// Wait until we've received a done signal for each worker.
	for _ = range done {
		fmt.Println("A worker has finished.")
	}
	fmt.Println("All workers have finished.")
}
