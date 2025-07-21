package main

import (
	"fmt"
	"errors"
)

func divide(a,b float64) (float64,error){
	if b==0 {
		return 0,errors.New("Cannot divide by zero")
	}
	return a/b,nil
}

func main(){
	results, err := divide(25,10)
	if err != nil {
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println("Results: ",results)
}