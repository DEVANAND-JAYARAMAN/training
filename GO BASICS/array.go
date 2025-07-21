package main
import "fmt"

func main(){
	var a int
	fmt.Println("Enter a number: ")
	fmt.Scan(&a)

	if a<0 {
		fmt.Println("Negative")
	} else if a>0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Zero")
	}
}