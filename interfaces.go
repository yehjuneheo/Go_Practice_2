package main

import (
	"fmt"
	"math"
)

// Shape is an interface that has a method signature for area
type Shape interface {
	Area() float64
}

// Rectangle is a struct with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle is a struct with radius
type Circle struct {
	Radius float64
}

// Implementing the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implementing the Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Function that takes a Shape interface as a parameter
func printArea(s Shape) {
	fmt.Printf("The area is: %.2f\n", s.Area())
}

func main4() {
	r := Rectangle{Width: 10.5, Height: 5.5}
	c := Circle{Radius: 10}

	printArea(r)
	printArea(c)
}
