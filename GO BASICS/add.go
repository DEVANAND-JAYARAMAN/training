//Add two integers and print the result

package main
import "fmt"
func main(){
	a:=100
	b:=9
	defer fmt.Println((a+b))
	defer fmt.Println(a-b)
	defer fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
	a+=100
	fmt.Println(a)
}