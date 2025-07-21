package main
import "fmt"

func main() {
    var str string
    fmt.Print("Enter string: ")
    fmt.Scanln(&str)

    seen := make(map[rune]bool)
    result:=""

    for _,ch := range str{
        if !seen[ch]{
            result+=string(ch)
            seen[ch]=true
        }
    }

    fmt.Println("After removing duplicates:", result)
}
