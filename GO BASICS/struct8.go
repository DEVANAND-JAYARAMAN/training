package main
import "fmt"
type student struct{
	name string
	marks float64
	rollno int
}

func(s student) function() {
	fmt.Println("Student Information: ")
	fmt.Println("Student name: ",s.name)
	fmt.Println("Student marks: ",s.marks)
	fmt.Println("Student rollno: ",s.rollno)
	
}

func main() {
	s1 := student{
		name : "JD",
		marks : 91,
		rollno: 8,
	}

	s1.function()
}