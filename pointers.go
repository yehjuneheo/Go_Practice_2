package main

import "fmt"

// Function that takes a pointer to an int and doubles the int's value
func double(n *int) {
	*n = *n * 2
}

func main10() {
	// Declare an int variable
	i := 10

	// Print the initial value of the int variable
	fmt.Println("Initial value of i:", i)

	// Get the address of the int variable and pass it to the double function
	double(&i)

	// Print the value of the int variable after calling the double function
	fmt.Println("Value of i after calling double:", i)
}
