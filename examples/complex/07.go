package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	const message string = "Hello, World!"
	var num1 int = 10
	var num2 int = 5

	sum := add(num1, num2)
	product := num1 * num2

	fmt.Println(message, "\n")
	fmt.Println("Sum:", sum, "\n")
	fmt.Println("Product:", product, "\n")
}
