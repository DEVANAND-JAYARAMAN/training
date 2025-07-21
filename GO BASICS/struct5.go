package main
import "fmt"

type Person struct{
	Name string
	Age int
	College string
	Salary int
}

func main(){
	var Pers1 Person
	var Pers2 Person
	Pers1.Name = "DEV ANAND"
	Pers1.Age = 21
	Pers1.College = "KIT"
	Pers1.Salary = 20000

	Pers2.Age = 22
	Pers2.Name = "Rakshith"
	Pers2.College = "KCE"
	Pers2.Salary = 25000

	fmt.Println(Pers1)
	fmt.Println(Pers2)
}