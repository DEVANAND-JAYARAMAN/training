package main
import "fmt"

func isPalindrome(str string) bool {
    n := len(str)
    for i := 0; i < n/2; i++ {
        if str[i] != str[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var str string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&str)

    if isPalindrome(str) {
        fmt.Println("It's a Palindrome")
    } else {
        fmt.Println("Not a Palindrome")
    }
}
