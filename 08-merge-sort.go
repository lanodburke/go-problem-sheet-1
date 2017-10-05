package main

import (
	"fmt"
	"sort"
)

func main(){

	var a, b [5]int
	
	fmt.Println("Please enter 5 numbers for 1st array")

	for i:=0; i<5;i++{
		fmt.Printf("Please enter number %d for array 1: ", i + 1)
		fmt.Scanf("%d ",&a[i])
		fmt.Printf("Please enter number %d for array 2: ", i + 1)
		fmt.Scanf("%d ",&b[i])
	}
	
	fmt.Printf("Sorted list: %v\n", sortlist(merge(a,b)))
}

func merge(a,b [5] int) [] int{
	x,y := a[:len(a)], b[:len(b)]
	result := append(x,y...)
	return result
}

func sortlist(a [] int) [] int {
	sort.Ints(a)
	return a
}