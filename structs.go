package main

import "fmt"

// Defining a struct type
type Person struct {
	Name string
	Age  int
}

// Function to print person details
func (p Person) printDetails() {
	fmt.Printf("Name: %s\nAge: %d\n", p.Name, p.Age)
}

func main2() {
	// Creating an instance of the Person struct
	p1 := Person{
		Name: "John Doe",
		Age:  30,
	}

	// Printing the details of the person
	p1.printDetails()
}
