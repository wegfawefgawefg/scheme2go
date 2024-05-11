package main

import (
	"fmt"
)

// Define a recursive function to compute factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Define a function using a map and array
func processItems() {
	scores := map[int]int{
		1: 90,
		2: 80,
		3: 70,
		5: 50,
		8: 30,
	}
	values := []int{1, 2, 3, 5, 8}
	for _, value := range values {
		if score, ok := scores[value]; ok {
			fmt.Println("Found:", score)
		} else {
			fmt.Println("Not found:", value)
		}
	}
}

// Use pattern matching to handle different types
func typeCheck(item interface{}) {
	switch v := item.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

// Pattern match with value
func valueCheck(item interface{}) {
	switch item {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown value")
	}
}

// Making arrays
func makeArray(anarray []int) {
	fmt.Println("Array vals:")
	for _, val := range anarray {
		fmt.Println(val)
	}
}

// Main function to call other functions
func main() {
	fmt.Println("This is the main function.")
	result := factorial(5)
	fmt.Println("Factorial calculated:", result)
	processItems()
	makeArray([]int{1, 2, 3, 5, 8})
	typeCheck(5)
	typeCheck("hello")
	valueCheck(1)
	valueCheck(4)
	fmt.Println("Main function completed.")
}
