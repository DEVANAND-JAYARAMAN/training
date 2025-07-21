package main
import ("fmt")
func main(){
	var a int=42
	var b *int=&a
	fmt.Println("Value of a: ",a)
	fmt.Println("Address of a: ",&a)
	fmt.Println("Address of a by b: ",b)
	fmt.Println("Value of a by b: ",*b)

	*b=1000
	fmt.Println("Value of a after changing: ",a)
	fmt.Println(a+10)
	
}