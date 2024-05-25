package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter Your First Name")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	name := strings.TrimSpace(input)
	fmt.Printf("Hello, %s! Welcome to the Evolza Team!\n", name)
}
