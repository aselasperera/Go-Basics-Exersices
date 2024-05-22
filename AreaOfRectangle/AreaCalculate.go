package main

import (
	"fmt"
)

func main() {
	// Ask the user for length and width
	var length, width float64
	fmt.Println("Enter the length and width of the rectangle:")
	fmt.Print("Length: ")
	fmt.Scanln(&length)
	fmt.Print("Width: ")
	fmt.Scanln(&width)

	// Calculate the area
	area := length * width

	// Print the result
	fmt.Printf("The area of the rectangle with length %.2f and width %.2f is %.2f\n", length, width, area)
}
