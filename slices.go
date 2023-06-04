package main

import "fmt"

func main7() {
	// Declare and initialize a slice
	s := []string{"apple", "banana", "cherry"}

	// Print the slice
	fmt.Println("Initial slice:", s)

	// Add an element to the slice
	s = append(s, "date")

	// Print the slice after adding an element
	fmt.Println("Slice after adding an element:", s)

	// Iterate over the slice and print each element
	for i, fruit := range s {
		fmt.Printf("Element at index %d: %s\n", i, fruit)
	}
}
