package main

import (
	"fmt"
)

func main6() {
	// Simple for loop
	fmt.Println("Simple for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While-like loop
	fmt.Println("While-like loop:")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// For-each loop
	fmt.Println("For-each loop:")
	nums := []int{1, 2, 3, 4, 5}
	for index, num := range nums {
		fmt.Printf("nums[%d] = %d\n", index, num)
	}

	// Infinite loop with break
	fmt.Println("Infinite loop with break:")
	i = 0
	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}
