package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	// Prompt the user for a password
	fmt.Println("Enter a password to check its strength:")
	reader := bufio.NewReader(os.Stdin)
	password, _ := reader.ReadString('\n')
	password = password[:len(password)-1] // Trim the newline character

	// Check the password strength
	length := len(password)
	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// Evaluate the password strength
	fmt.Println("Password strength evaluation:")
	if length >= 8 {
		fmt.Println("- Length: Pass")
	} else {
		fmt.Println("- Length: Fail (must be at least 8 characters)")
	}

	if hasUpper {
		fmt.Println("- Uppercase: Pass")
	} else {
		fmt.Println("- Uppercase: Fail (must include at least one uppercase letter)")
	}

	if hasLower {
		fmt.Println("- Lowercase: Pass")
	} else {
		fmt.Println("- Lowercase: Fail (must include at least one lowercase letter)")
	}

	if hasDigit {
		fmt.Println("- Digit: Pass")
	} else {
		fmt.Println("- Digit: Fail (must include at least one digit)")
	}

	if hasSpecial {
		fmt.Println("- Special Character: Pass")
	} else {
		fmt.Println("- Special Character: Fail (must include at least one special character)")
	}

	// Determine overall strength
	if length >= 8 && hasUpper && hasLower && hasDigit && hasSpecial {
		fmt.Println("Overall Strength: Strong")
	} else if length >= 6 && hasUpper && hasLower && (hasDigit || hasSpecial) {
		fmt.Println("Overall Strength: Moderate")
	} else {
		fmt.Println("Overall Strength: Weak")
	}
}
