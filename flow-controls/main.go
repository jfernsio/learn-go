package main

import (
	"fmt"
)

func main() {
	fmt.Println("BAsic for loop")
	    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

fmt.Println(" for as while loop")
	//for as while 
	    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

	fmt.Println("infinre loop")
	//infinte loop
    j := 0
    for {
        if j >= 5 {
            break
        }
        fmt.Println(j)
        j++
    }

	fmt.Println("range based for loop")
	//rnage for loop
   numbers := []int{10, 20, 30}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

	//basic switch
	day := 2
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Another day")
	}

	fmt.Println("mutilpe case swirch")
	//mutiple cases switch
	switch day {
	case 1, 2, 3:
		fmt.Println("Start of the week")
	case 4, 5:
		fmt.Println("Midweek")
	default:
		fmt.Println("weekedn")
	}

	var x interface{} = "abc"
	switch x.(type) {
	case int:
		fmt.Println("int tyope")
	case string:
		fmt.Println("string typw")
	default:
		fmt.Println("unkown type")
	}
}
