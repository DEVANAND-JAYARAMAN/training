package main
import "fmt"

type class struct{
	name string
	rollno int
	cgpa float64
}

func main(){
	class1:=class{
		name:"Dev Anand",
		rollno:07,
		cgpa:7.32}
			fmt.Println(class1)

	fmt.Println(class1.name)
	fmt.Println(class1.rollno)
	fmt.Println(class1.cgpa)
	fmt.Println("---------------------------")
}

