//Count how many even and odd numbers are in an array.
package main
import "fmt"
func main(){
	var odd,even int=0,0
	fmt.Println("Enter length of an array: ")
	var a int
	fmt.Scan(&a)

	var arr=[]int{}
	for i:=0;i<len(a);i++ {
		if(arr[i]%2==0){
			even++
		} else{
			odd++
		}
	}
	fmt.Println("Odd: ",odd,"Even: ",even)
}