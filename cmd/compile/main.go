package main

import (
	"fmt"
	"os"
	"scheme2go/pkg/compile"
)

func main() {
	inputPath := "./cmd/compile/code2.rkt"
	input, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	strInput := string(input)

	// print the full program
	fmt.Println("Code:")
	fmt.Println(strInput)
	fmt.Println()

	output, err := compile.Compile(strInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// print the compiled program
	fmt.Println("Compiled:")
	fmt.Println(output)
}

// func add(a int, b int) int {
// 	return a + b
// }
