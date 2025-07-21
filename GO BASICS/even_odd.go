//✅ 4. Check if a number is even or odd

package main
import "fmt"

func main(){
	var num int
	fmt.Scan(&num);
	if(num%2==0){
		fmt.Println("Even")
	} else{
		fmt.Println("Odd")
	}
}