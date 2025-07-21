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

type Square struct{
	side float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func (s Square) Area() float64{
	return math.Pow(s.side,2)
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

func PrintArea(sh Shape){
	fmt.Println("Area is: ",sh.Area())
}

func main(){
	r:=Rectangle{4,5}
	s:=Square{10}
	c:=Circle{5}
	PrintArea(r)
	PrintArea(s)
	PrintArea(c)
	
}