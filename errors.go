package main

import (
	"errors"
	"fmt"
)

// Function that divides two numbers and returns an error if the divisor is zero
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, nil
}

func main5() {
	// Successful division
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 divided by 2 is", result)
	}

	// Division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 divided by 0 is", result)
	}
}
