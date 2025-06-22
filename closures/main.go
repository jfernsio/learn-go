package main

import (
	"fmt"
)

// a closure is a function that has access to its own scope, the scope of its outer functions, and the global scope
func adder() func(int) int {
	x := 5
	return func(y int) int {
		return x + y
	}
}

func counter() (func() int, func() int) {
	x := 0
	increment := func() int {
		x++
		return x
	}
	decrement := func() int {
		x--
		return x
	}
	return increment, decrement
}

func main() {
	add := adder()
	fmt.Println(add(3))

	inc, dec := counter()
	fmt.Println("inc",inc())
	fmt.Println("inc",inc())
	fmt.Println("inc",inc())

	fmt.Println("decc",dec())
	fmt.Println("decc",dec())

}
