package main
import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

type Rectangle struct{
	width, height float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) Area() float64{
	return r.height*r.width
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func main(){
	rect := Rectangle{
		width: 5,height: 4}
	cir := Circle{
		radius:2,
	}
	// fmt.Println(rect.Area())
	// fmt.Println(cir.Area())
	fmt.Println(CalculateArea(rect))
	fmt.Println(CalculateArea(cir))
}

func CalculateArea(s Shape) float64{
	return s.Area()
}