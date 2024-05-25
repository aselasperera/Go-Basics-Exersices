package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1

	// Initialize the guess variable
	var guess int

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess what it is?")

	// Loop until the user guesses the correct number
	for {
		// Prompt the user for a guess
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check the user's guess
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
