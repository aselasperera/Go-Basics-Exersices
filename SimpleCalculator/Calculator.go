package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Ask the user for numbers and operator
	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)
	fmt.Println("Enter the operator (+, -, *, /):")
	fmt.Scanln(&operator)

	// Perform calculation based on the operator
	result := calculate(num1, num2, operator)

	// Print the result
	fmt.Printf("Result: %.2f\n", result)
}

func calculate(num1, num2 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Error: Invalid operator")
	}

	return result
}
