package main
import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	sum := 0
	for num != 0 {
		d := num % 10
		sum += d
		num /= 10
	}

	fmt.Println("Sum of digits:", sum)
}
