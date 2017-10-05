// source: https://gist.github.com/itang/8109481

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(reverse("Hello There!"))
}

func reverse(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}