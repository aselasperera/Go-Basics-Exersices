package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	inputFile := "input.txt"
	inputData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	lines := strings.Split(string(inputData), "\n")
	lineCount := len(lines)
	wordCount := 0
	for _, line := range lines {
		words := strings.Fields(line)
		wordCount += len(words)
	}

	outputFile := "output.txt"
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer output.Close()

	output.WriteString(fmt.Sprintf("Number of lines: %d\n", lineCount))
	output.WriteString(fmt.Sprintf("Number of words: %d\n", wordCount))

	fmt.Println("Data processed and written to output.txt")
}
