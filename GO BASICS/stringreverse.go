package main
import "fmt"

func main() {
    var str string
    fmt.Print("Enter string: ")
    fmt.Scanln(&str)

    reversed := ""
    for _, ch := range str {
        reversed = string(ch) + reversed
    }

    fmt.Println("Reversed:", reversed)
}
