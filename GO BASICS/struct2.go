package main
import "fmt"
type car struct{
	Brand string
	Year int
	Model string
}

func main(){
	myCar:=car{
		Brand:"Volkswagen Polo",
		Model:"GT",
	}
	fmt.Println(myCar.Brand)
	fmt.Println(myCar.Year)
	fmt.Println(myCar.Model)

	
}
