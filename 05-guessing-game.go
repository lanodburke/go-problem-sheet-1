// source: http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true //integer is in list
        }
    }
    return false
}

func main() {
    myrand := xrand(1, 100)
    tries := 0
	var guess int
	var guesses [] int 

    fmt.Println("Welcome to Guess My Number Game!")
    for guess != myrand {
        fmt.Println("Guess a number between 1 and 100")
        fmt.Scanf("%v", &guess)
        tries++
        if guess > myrand {
            fmt.Println("Too high")
        } else if guess < myrand {
            fmt.Println("Too low")
        } else {
            fmt.Printf("Good job! You guessed it in %v tries\n", tries)
            break
		}
		
		if contains(guesses, guess) == true{
            fmt.Printf("That number has already been entered\n")
            tries--
        } else {
            //add guess to slice of guessed numbers
            guesses = append(guesses, guess)
            
        }

    }
}
