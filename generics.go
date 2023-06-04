// This is not valid Go code as of September 2021
package main

import "fmt"

type List[T any] []T

func PrintItems[T any](list List[T]) {
	for _, item := range list {
		fmt.Println(item)
	}
}

func main13() {
	integers := List[int]{1, 2, 3, 4, 5}
	PrintItems[int](integers)

	strings := List[string]{"hello", "world"}
	PrintItems[string](strings)
}
