package main

import (
	"fmt"
	"os"
)

func main(){
	file, err := os.Create("sample.txt")
	if err != nil{
		fmt.Println("Error: ",err)
		return
	}
	defer file.Close()

	file.WriteString("Welcome to the Go File Operations")
	fmt.Println("File was created and the content was written on it")
}