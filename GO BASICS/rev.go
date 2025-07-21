package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    reverse := 0

    for num != 0 {
        digit := num % 10             
        reverse = reverse*10 + digit  
        num = num / 10                
    }
	
    fmt.Println("Reversed number:", reverse)
}
