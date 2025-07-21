package main
import (
	"fmt"
)

func main(){
	var a int
	fmt.Scan(&a)
	for i:=0;i<a;i++ {
		if(i%3 == 0 && i%4 == 0){
			fmt.Print(i,"  ")
		} else{
			continue
		}
	}
}