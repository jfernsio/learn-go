package main

import (
	"fmt"
)

func main() {
	var a []int = []int{1, 2, 3}
	fmt.Println("a:", a)
	nums := []int{10, 20, 30}
	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))

	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]

	fmt.Println("slice:", slice)
	fmt.Println("len:", len(slice)) // 3 (elements shown)
	fmt.Println("cap:", cap(slice)) // 4 (from index 1 to end of array)

	s := make([]int, 2, 5)      // length = 2, capacity = 5
	fmt.Println("slice:", s)    // [0 0]
	fmt.Println("len:", len(s)) // 2
	fmt.Println("cap:", cap(s)) // 5

	s = append(s, 1, 2, 3)          // now has 5 elements
	fmt.Println("after append:", s) // [0 0 1 2 3]
	fmt.Println("len:", len(s))     // 5
	fmt.Println("cap:", cap(s))     // 5

	s = append(s, 99)            // capacity exceeded!
	fmt.Println("new array:", s) // [0 0 1 2 3 99]
	fmt.Println("len:", len(s))  // 6
	fmt.Println("cap:", cap(s))  // likely 10 (Go doubles it)
}
