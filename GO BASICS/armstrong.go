//Check if a number is an Armstrong number

package main
import(
	"fmt"
	"strconv"
	"math"
)
func main(){
	var a int
	fmt.Scan(&a)
	original:=a
	leng:=len(strconv.Itoa(a));
	sum:=0.0

	for a!=0{
		digit:=a%10;
		sum=(math.Pow(float64(digit),float64(leng)))+sum
		a/=10
	}
	if int(sum) == original {
		fmt.Printf("%d is an Armstrong number.\n", original)
	} else {
		fmt.Printf("%d is not an Armstrong number.\n", original)
	}
}