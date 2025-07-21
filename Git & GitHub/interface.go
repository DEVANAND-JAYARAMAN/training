package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	shapes := []Shape{
		Circle{Radius: 4},
		Circle{Radius: 3.5},
		Rectangle{Width: 6, Height: 7},
		Rectangle{Width: 6.5, Height: 7.5},
		
		
	}
	for i, shape := range shapes {
		fmt.Printf("Shape %d area: %.2f\n", i+1, shape.Area())
	}
}
