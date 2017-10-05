// source https://gist.github.com/sighmin/9173219

package main 

import (
	"fmt"
	"math"
)

func main(){
	var num float64

	fmt.Println("Enter number to square root: ")
	fmt.Scanf("%f", &num)

	fmt.Println("Newton's method: ", Newt(num))
	fmt.Println("Math.Sqrt method: ", math.Sqrt(num))
}

func Newt(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}