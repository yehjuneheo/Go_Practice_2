package main

import (
	"fmt"
)

func main() {
	// Integer variable
	var i int = 10
	fmt.Printf("Integer: %d\n", i)

	// Float variable
	var f float64 = 3.1415
	fmt.Printf("Float: %f\n", f)

	// Boolean variable
	var b bool = true
	fmt.Printf("Boolean: %t\n", b)

	// String variable
	var s string = "Hello, World!"
	fmt.Printf("String: %s\n", s)

	// Array variable
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Array: %v\n", arr)

	// Slice variable
	var sl []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", sl)

	// Map variable
	var m map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("Map: %v\n", m)

	// Pointer variable
	var p *int = &i
	fmt.Printf("Pointer: %p\n", p)
}