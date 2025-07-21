package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Largest number is:", a)
	} else if b > a {
		fmt.Println("Largest number is:", b)
	} else {
		fmt.Println("Both numbers are equal.")
	}
}
