package main
import "fmt"

func main() {
    var city string
    fmt.Print("Enter your city: ")
    fmt.Scanln(&city)
    fmt.Println("You live in", city)
}
