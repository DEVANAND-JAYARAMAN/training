package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your address: ")
	fulladdr,_ := reader.ReadString('\n')
	fmt.Println("Your address is: "+fulladdr)
}