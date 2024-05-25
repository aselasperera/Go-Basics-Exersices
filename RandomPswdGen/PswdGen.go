package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:'\",.<>?/\\`~"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for password length
	fmt.Println("Enter the desired password length:")
	var length int
	fmt.Scanf("%d\n", &length)

	// Prompt the user for character types
	fmt.Println("Include lowercase letters? (y/n)")
	includeLowercase := getYesNoInput(reader)
	fmt.Println("Include uppercase letters? (y/n)")
	includeUppercase := getYesNoInput(reader)
	fmt.Println("Include digits? (y/n)")
	includeDigits := getYesNoInput(reader)
	fmt.Println("Include special characters? (y/n)")
	includeSpecialChars := getYesNoInput(reader)

	// Generate the password
	password := generatePassword(length, includeLowercase, includeUppercase, includeDigits, includeSpecialChars)

	// Print the generated password
	fmt.Printf("Generated password: %s\n", password)
}

func getYesNoInput(reader *bufio.Reader) bool {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strings.ToLower(input) == "y"
}

func generatePassword(length int, includeLowercase, includeUppercase, includeDigits, includeSpecialChars bool) string {
	var charPool string

	if includeLowercase {
		charPool += lowercaseLetters
	}
	if includeUppercase {
		charPool += uppercaseLetters
	}
	if includeDigits {
		charPool += digits
	}
	if includeSpecialChars {
		charPool += specialChars
	}

	if len(charPool) == 0 {
		fmt.Println("Error: At least one character type must be selected")
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charPool[rand.Intn(len(charPool))]
	}

	return string(password)
}
