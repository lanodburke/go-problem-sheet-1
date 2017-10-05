// source: https://github.com/chirag04/golang-project_euler/blob/master/p020.go

package main

import "fmt"

func main(){
	sum := factorial_sum()
	fmt.Printf("The sum of 100!: %d\n", sum)
}

func factorial_sum() int {
    sum := 0;
    digits := [200]int{};
    digits[0] = 1;
    for i := 2; i <= 100; i++ {
    	for j := 0; j < len(digits); j++ {
    		digits[j] *= i;
    		if j > 0 && digits[j - 1] > 9 {
    			digits[j] += int(digits[j - 1] / 10);
    			digits[j - 1] %= 10;
    		}
    	}
    }
    for i := 0; i < len(digits); i++ {
    	sum += digits[i];
    }
    return sum;
}