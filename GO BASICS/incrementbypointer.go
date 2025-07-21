package main
import "fmt"
func main(){
	x:=10
	increment(&x)
	fmt.Println(x)
}
func increment(val *int){
	*val=*val+7
}