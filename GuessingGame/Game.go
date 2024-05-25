package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	var guess int

	fmt.Println("Welcome to the Evolza Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
	fmt.Println("Let's get started!")

	for {

		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d correctly.\n", target)
			break
		}
	}
}
