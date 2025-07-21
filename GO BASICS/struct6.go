package main
import "fmt"

type car struct{
	Brand string
	Model string
	Price int
}

func main(){
	favcar := car{"BMW","x7",25000000}
	
	var jdcar car
	jdcar.Brand = "Porsche"
	jdcar.Model = "GT3RS"
	jdcar.Price = 15000000

	supercar := car{
		Brand: "Volkswagen",
		Model: "Virtus GT",
		Price : 3500000,
	}

	fmt.Println(supercar)
	fmt.Println(jdcar)
	fmt.Println(favcar)
}