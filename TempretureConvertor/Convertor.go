package main

import (
	"fmt"
)

func main() {
	var option string
	var heat, convertedTemp float64

	fmt.Println("Temperature Converter")
	fmt.Println("1. Convert Celsius to Fahrenheit")
	fmt.Println("2. Convert Fahrenheit to Celsius")
	fmt.Print("Enter your choice (1 or 2): ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Print("Enter temperature in Celsius: ")
		fmt.Scanln(&heat)
		convertedTemp = celsiusToFahrenheit(heat)
		fmt.Printf("Temperature in Fahrenheit: %.2f\n", convertedTemp)
	case "2":
		fmt.Print("Enter temperature in Fahrenheit: ")
		fmt.Scanln(&heat)
		convertedTemp = fahrenheitToCelsius(heat)
		fmt.Printf("Temperature in Celsius: %.2f\n", convertedTemp)
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
