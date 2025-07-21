//reverse the array
package main
import "fmt"

func main() {
    arr := [5]int{10, 20, 30, 40, 50}

    fmt.Println("Reversed array:")
    for i := len(arr) - 1; i >= 0; i-- {
        fmt.Println(arr[i])
    }
}
