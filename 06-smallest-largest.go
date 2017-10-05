// source: https://gist.github.com/pyk/10519339

package main

import "fmt"

func main() {
  var n, smallest, biggest int
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }

  for _,v:=range x {
    if v>n {
      fmt.Println(v,">",n)
      n = v
      biggest = n
    } else {
      fmt.Println(v,"<",n)
    }
  }

  fmt.Println("The biggest number is ", biggest)
  for _,v:=range x {
    if v>n {
      fmt.Println(v,">",n)
    } else {
      fmt.Println(v,"<",n)
      n = v
      smallest = n
    }
  }
  fmt.Println("The smallest number is ", smallest)
}