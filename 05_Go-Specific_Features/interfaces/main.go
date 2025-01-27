package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.length
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.length)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	rect := Rectangle{Width: 5, length: 10}

	var shape Shape = rect

	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())

	circle := Circle{Radius: 5}
	shape = circle

	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
}
