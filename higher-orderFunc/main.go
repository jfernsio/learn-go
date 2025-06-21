package main

import "fmt"

// passing func as argument
// Higher-order function
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Functions to pass
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

// returning a func
func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}
func main() {
	// Using the higher-order function
	fmt.Println("Sum:", operate(3, 5, add))
	fmt.Println("Product:", operate(3, 5, multiply))

	double := multiplier(2) // Returns a function to double the input
	triple := multiplier(3) // Returns a function to triple the input

	fmt.Println("Double of 5:", double(5)) 
	fmt.Println("Triple of 5:", triple(5))
}
