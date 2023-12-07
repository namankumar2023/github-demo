package main

import "fmt"

func main() {
	res := add(2, 5)
	fmt.Println("Result: ", res)
	fmt.Println("End Of Program")
}

// method is for adding two numbers
func add(num1 int, num2 int) int {
	return num1 + num2
}
