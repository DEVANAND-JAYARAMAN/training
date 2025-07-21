package main
import(
	"fmt"
)
func main(){
	nums := []string{}
	nums=append(nums, "JD","Dev","Anand","Anusha","Dev Anand Anusha")
	fmt.Println("Array: ",nums)
	for i,name := range nums{
		fmt.Printf("Index of %d: %v\n",i,name)
	}
	for i,name := range nums{
		if name=="Anusha"{
			nums[i]="ANUSHA"
		}
	}
	fmt.Println("after update: ",nums)
}