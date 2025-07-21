package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Please enter 4 arguments in this order: program_name.go num1 operator num2")
		return
	}

	str1 := os.Args[1]
	operator := os.Args[2]
	str2 := os.Args[3]

	num1, err1 := strconv.ParseFloat(str1, 64)
	num2, err2 := strconv.ParseFloat(str2, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers.")
		return
	}
	var result float64
	var err error
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Errorf("Error: Division by zero is not allowed")
			return
		} else {
			result = num1 / num2
		}
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result is: ", result)
	}

}
