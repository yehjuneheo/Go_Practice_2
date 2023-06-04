package main

import "fmt"

// Variadic function to calculate the sum of an arbitrary number of values
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func main9() {
	// Call the variadic function
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

	// Anonymous function
	double := func(n int) int {
		return n * 2
	}
	fmt.Println("Double of 5:", double(5))

	// Function as a first-class citizen
	funcs := map[string]func(int) int{
		"double": double,
		"square": func(n int) int { return n * n },
	}
	fmt.Println("Double of 6:", funcs["double"](6))
	fmt.Println("Square of 3:", funcs["square"](3))
}
