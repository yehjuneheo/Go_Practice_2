package main

import (
	"fmt"
)

// Function with no parameters
func printMessage() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func printCustomMessage(message string) {
	fmt.Println(message)
}

// Function that returns a value
func square(n int) int {
	return n * n
}

// Function that returns multiple values
func getCoordinates() (int, int) {
	x := 10
	y := 20
	return x, y
}

func main3() {
	// Call function without parameters
	printMessage()

	// Call function with parameters
	printCustomMessage("Hello, Go!")

	// Call function that returns a value
	fmt.Printf("Square of 5: %d\n", square(5))

	// Call function that returns multiple values
	x, y := getCoordinates()
	fmt.Printf("Coordinates: %d, %d\n", x, y)
}
