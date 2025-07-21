package main
import "fmt"

func main(){
	var a = [5] int {1,2,3,4,5}
	for index,_ := range a{
		fmt.Printf("%d\n",index)
	}
}