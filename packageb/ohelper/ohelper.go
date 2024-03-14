package ohelper

import (
	"fmt"
)

// callee
func PrintInput(input string) string {
	// ruleid: test_inter_analysis
	fmt.Println("Output:", input)
	return "Processed: " + input
}
