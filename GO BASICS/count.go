//✅ 4. Count Number of Digits

package main
import (
	"fmt"
)
func main(){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	count:=0
	for num!=0{
		num/=10
		count++
	}
	fmt.Print("Digits in number is: ",count)
}
