package main

import "fmt"

func main() {
	fmt.Println("Hello")
	res := add(2, 5)
	fmt.Println("Result: ", res)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
