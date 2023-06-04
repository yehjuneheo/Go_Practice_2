package main

import "fmt"

func main8() {
	// Declare and initialize a map
	m := map[string]int{"apple": 5, "banana": 10, "cherry": 15}

	// Print the map
	fmt.Println("Initial map:", m)

	// Add an element to the map
	m["date"] = 20

	// Print the map after adding an element
	fmt.Println("Map after adding an element:", m)

	// Delete an element from the map
	delete(m, "apple")

	// Print the map after deleting an element
	fmt.Println("Map after deleting an element:", m)

	// Iterate over the map and print each key and value
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
