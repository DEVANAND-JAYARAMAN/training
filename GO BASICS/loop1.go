package main
import "fmt"
func main(){
	var fruits=[3]string{"Apple","Banana","Mango"}
	fmt.Println(len(fruits))

	for _,ind := range fruits{
		fmt.Println(ind)
	}
	for i:=0;i<len(fruits);i++{
		fmt.Println(fruits[i])

	}
}