package main

import "fmt"

func main () {
	//decakring a pointer
	var z *int //--> ptr to an int

	//assigning a pointer
	x := 10 
	z = &x //z now holds the address of x

	//dereferencing a pointer
	fmt.Println("addres of x :", z)
	fmt.Println("value of x :", *z)

	//ptr to ptr 
	pp := &z
	fmt.Println("address of z:", pp)
	fmt.Println("value of z:", **pp)
	
}
