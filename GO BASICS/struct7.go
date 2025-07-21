package main
import "fmt"

type school struct{
	name string
	marks int
	class int
}

func main(){
	// var student1 school
	// student1.name = "Dev Anand J"
	// student1.marks = 98
	// student1.class = 10

	student1 := school{
		name : "JD",
		marks: 100,
		class: 12,
		
	}
	student1.name = "Dev Anand J"

	fmt.Println(student1)
}