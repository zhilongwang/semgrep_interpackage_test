package main

import (
	"bufio"
	"fmt"
	"os"

	"example/helper"
	"pkgexample/ohelper"
)

func main() {
	// Read input from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	input, _ := reader.ReadString('\n')

	// Pass input to callee
	output := processInput(input)

	helper.PrintInput(input)
	ohelper.PrintInput(input)

	// ruleid: test_inter_analysis
	fmt.Println("Output:", output)
}

// callee
func processInput(input string) string {
	// ruleid: test_inter_analysis
	fmt.Println("Output:", input)
	return "Processed: " + input
}
