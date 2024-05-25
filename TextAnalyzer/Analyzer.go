package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Enter a text string:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	wordQuantity := countWords(input)

	charQuantity := countCharacters(input)

	fmt.Println("Enter a specific letter to count:")
	letter, _ := reader.ReadString('\n')
	letter = strings.TrimSpace(letter)

	letterQuantity := countSpecificLetter(input, letter)

	fmt.Printf("Number of words: %d\n", wordQuantity)
	fmt.Printf("Number of characters (excluding spaces): %d\n", charQuantity)
	fmt.Printf("Occurrences of '%s': %d\n", letter, letterQuantity)
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func countCharacters(text string) int {
	count := 0
	for _, char := range text {
		if !unicode.IsSpace(char) {
			count++
		}
	}
	return count
}

func countSpecificLetter(text, letter string) int {
	count := 0
	for _, char := range text {
		if string(char) == letter {
			count++
		}
	}
	return count
}
