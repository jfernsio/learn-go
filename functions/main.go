package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func name() {
	fmt.Println("golang")
}
func namewparams(name string) string {
	return "Hello, " + name
}

// varadic func
func add(a ...int) int {
	total := 0
	for i := range a {
		total += a[i]
	}
	return total
}

// func closure
func incrementer() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// func w mutiple return vals
func convert(str string) (int, error) {
	return strconv.Atoi(str)
}

// func w named val returns
func Area(x, y float64) (area float64, p float64) {
	area = x * y
	p = 2 * (x + y)
	return

}


//func making a http req
func getRequest (url string) (*http.Response,error) {
	res, err := http.Get(url)
	return res,err
}


func main() {
	name()
	fmt.Println(namewparams("World"))
	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println(add(10, 20, 30, 40, 50, 60, 70, 80, 90, 100))
	incr := incrementer()

	fmt.Println(incr()) // Output: 1
	fmt.Println(incr()) // Output: 2

	i, e := convert("123")
	if e != nil {
		fmt.Println("Error converting string to int:", e)
		return
	} else {
		fmt.Println("Converted value:", i)
	}

	//w wrong val
	// i, e := convert("abc")
	// if e != nil {
	// 	fmt.Println("Error converting string to int:", e)

	// 	return
	// } else {
	// 	fmt.Println("Converted value:", i)
	// }

	fmt.Println(Area(5, 10))

	    response, err := getRequest("https://cas.com")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Response Status:", response.Status)
    }
}
