// source: http://www.golangpro.com/2016/01/check-string-palindrome-go.html

package main

import (
 "fmt"
 "strings"
)

func main() {

 var palindrome string
 fmt.Println("Enter string:")
 fmt.Scanf("%s\n", &palindrome)
 palindrome = strings.ToLower(palindrome)
 fmt.Println(isPalindrome(palindrome))
}
//Function to test if the string entered is a Palindrome

func isPalindrome(s string) string {
 mid := len(s) / 2
 last := len(s) - 1
 for i := 0; i < mid; i++ {
  if s[i] != s[last-i] {
   return "No, its not a palindrome"
  }
 }
 return "Yes! You've entered a Palindrome"
}