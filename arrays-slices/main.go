package main

import "fmt"

func main() {
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	fmt.Println(arr1[1])

	//decalre n intial arr
	arr := [3]int{1, 2, 3}

	arr[0] = 10 // Set the first element to 10
	fmt.Println(arr[0])

	// mutidimensional arr
	var matrix [2][3]int // A 2x3 matrix (2 rows, 3 columns)
	matrix[0][0] = 1
	fmt.Println(matrix[0][0])

	//slice
	// var slice []int //declares a slice

	// slice also can be created from an arraay
	arrr := [3]int{1, 2, 3}

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // --> creates a slice which includes elements 1 throuh 3
	//includes the first element, but excludes the last one
	fmt.Println("primes", s)

	entierArray := arrr[:] // ==>creates slice of whole arr
	fmt.Println("enterier arr slice", entierArray)

	//creating slices w make func
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)

	// apppending to silce
	slice := []int{1, 2, 3}

	slice = append(slice, 4, 5) // Adds 4 and 5 to the slice
	fmt.Println(slice)

	//range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
