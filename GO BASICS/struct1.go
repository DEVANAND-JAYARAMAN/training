package main
import "fmt"

type Car struct{
	Brand string
	Year int
}
func main(){
	mycar:=Car{"Ford",1969}
	fmt.Println(mycar)
	fmt.Println("Brand:", mycar.Brand)
    fmt.Println("Year:", mycar.Year)
	dev()
}
func dev(){
	mycar:=Car{"BMW",2004}
	myCar:=Car{"Audi",2005}
	
	fmt.Println(mycar)
	fmt.Println(myCar)
	
	fmt.Println("Brand:", mycar.Brand)
    fmt.Println("Year:", mycar.Year)

	fmt.Println("Brand:", myCar.Brand)
    fmt.Println("Year:", myCar.Year)
}

