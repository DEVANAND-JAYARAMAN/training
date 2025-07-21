//using pointer, swapping two numbers using temp variable

package main
import "fmt"

func main(){
	a:=10
	b:=20
	swap(&a,&b)
	fmt.Print(a)
	fmt.Print(b)
}

func swap(x,y *int){
	temp:=*x
	*x=*y
	*y=temp
}