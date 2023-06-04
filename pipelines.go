package main

import (
	"fmt"
)

// generator() function that converts a list of integers to a channel that emits the integers in the list.
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// square() function that receives integers, squares them, and sends them on a different channel.
func square2(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// function to compute a square of a number
func square3(n int) int {
	return n * n
}

func main16() {
	// Set up the pipeline.
	c := generator(2, 3, 4, 5)
	out := square2(c)

	// Consume the output.
	for n := range out {
		fmt.Println(n)
	}
}
