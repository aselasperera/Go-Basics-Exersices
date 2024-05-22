package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ask the user for their name
	fmt.Print("Enter your name: ")

	// Read user input from standard input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim whitespace from the input and convert to title case
	name := strings.TrimSpace(input)

	// Greet the user
	fmt.Printf("Hello, %s!\n", name)
}
