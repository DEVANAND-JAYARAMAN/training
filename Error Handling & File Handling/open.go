package main

import (
	"log"
	"os"
)

func main(){
	file, err := os.Open("sample.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
}