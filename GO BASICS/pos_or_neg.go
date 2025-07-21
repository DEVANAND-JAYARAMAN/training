//Check if a number is positive or negative

package main
import "fmt"

func main(){
	var num int=-1000000000000000
	if(num>0){
		fmt.Println("Positive")
	} else if(num<0){
		fmt.Println("Negative")
	} else{
		fmt.Println("0")
	}
}