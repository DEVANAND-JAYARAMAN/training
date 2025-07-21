package main

import "fmt"

type rectangle struct {
	Width  float64
	Height float64
}

type circle struct {
	Radius float64
}

func (rec rectangle) AreaR() float64 {
	return rec.Width * rec.Height
}

func (c circle) AreaC() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	rect := rectangle{
		Width:  9,
		Height: 9,
	}
	fmt.Printf("Area of the rectangle is: %.2f\n", rect.AreaR())
	fmt.Println("----------------------")

	var cir circle
	cir.Radius = 6
	fmt.Printf("Area of the circle is: %.2f\n", cir.AreaC())
	fmt.Println("----------------------")

	people := []struct {
		Name string
		Age  int
	}{
		{"Dev Anand", 20},
		{"Anusha", 21},
		{"Dev Anand", 20},
		{"Anusha", 21},
		{"Dev Anand", 20},
		{"Anusha", 21},
		{"Dev Anand", 20},
		{"Anusha", 21},
		{"Dev Anand", 20},
		{"Anusha", 21},
		{"Dev Anand", 20},
		{"Anusha", 21},
	}

	for _, person := range people {
		fmt.Println(person.Name,"\t", person.Age)
		fmt.Println("----------------------")

	}
}
