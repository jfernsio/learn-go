package main

import (
	// "golang.org/x/tour/wc"
	"fmt"
	"strings"
)
func wordCounts (s string) map[string]int {
	words := make(map[string]int)
	counts := strings.Fields(s)

	for _, word := range counts {
	return 	words[word]++
}
}
func main() {
	var m map[string]int //decalres a nil map w string keys and int vals
	fmt.Println("m:", m)

	m0 := make(map[string]string) // --> empty map
	fmt.Println("m0:", m0)

	m1 := make(map[string]int, 5) // ==> map w capcity of 5
	fmt.Println("m1:", m1)

	//also can decalre map w inital vals
	m2 := map[string]int{  //--> aslo a map literal
		"one": 1,
		"two": 2,
	}
	fmt.Println("m2:", m2)

	//add a key val pair
	m0["nmae"] = "john ode"
	fmt.Println("m0 after adding key:", m0)

	//updte key val
	m2["one"] = 11
	fmt.Println("m2 after updating key:", m2)

	//check if ke exits
	val, ok := m2["two"]
	if ok {
		fmt.Println("key 'two' exists with value:", val)
	} else {
		fmt.Println("key 'two' does not exist")
	}
	//also
	fmt.Println("The value:", val, "Present?", ok)

	//delete a key
	delete(m2, "one")
	fmt.Println("m2 after deleting key 'one':", m2)

	//clear a map
	clear(m0)
	fmt.Println("m0 after clearing:", m0)

	// multideimensonal maps

wordCounts("hello hi hello")
}