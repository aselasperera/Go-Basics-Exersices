package main

import (
	"fmt"
	"math"
)

func main() {
	// Input values
	principal := 10000.0 // Principal amount
	interestRate := 0.05 // Interest rate (5%)
	loanTerm := 12       // Loan term in months

	// Calculate monthly interest rate
	monthlyInterestRate := interestRate / 12

	// Calculate monthly payment
	monthlyPayment := (principal * monthlyInterestRate) / (1 - math.Pow(1+monthlyInterestRate, float64(-loanTerm)))

	// Print the monthly payment
	fmt.Printf("Monthly Payment: $%.2f\n", monthlyPayment)
}
