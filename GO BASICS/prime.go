package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)

	isPrime := true
	if num < 1 {
		isPrime = false
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not a Prime number")
	}

}
