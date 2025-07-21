package main
import "fmt"

func main(){
	x:=10
	change(&x)
	fmt.Println(x)

	p := new(int) // allocates memory for an int
	*p = 50
	fmt.Println(p) // 50



}
func change(value *int){
	*value=10000000
}