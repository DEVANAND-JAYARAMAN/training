package main
import "fmt"

func main() {
    nums := []int{10, 20, 30, 40, 50}
    fmt.Println("Original:", nums)
    indexToDelete := 2 
    nums = append(nums[:indexToDelete], nums[indexToDelete+1:]...)
    fmt.Println("After delete:", nums)
}
