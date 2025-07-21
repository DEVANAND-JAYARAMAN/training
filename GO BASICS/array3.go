//Find the maximum element in an array.

package main
import "fmt"
func main(){
	var arr=[3]int{}
	for i:=0;i<len(arr);i++{
		fmt.Scan(&arr[i])
	}
	max:=0
	for i:=0;i<len(arr);i++{
		if(arr[i]>=max){
			max=arr[i];
		}
	}
	fmt.Println("Maximum is:",max)	
}