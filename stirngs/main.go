package main

import (
	"fmt"
)

func main() {
	var greeting string = "Hello, Gophers!"
	rawString := `This is a "raw" 
string\n` //-->  Represent raw strings, where escape sequences are treated literally.

	fmt.Println(greeting)
	fmt.Println(rawString)

	// length of string in bytes
	fmt.Println("legth of greeting:", len(greeting))

	// concatenation
	greeting += "dont leave me high and dry"
	fmt.Println("greeting after concatenation:", greeting)

	// accesinff chars as byte
	firstByte := greeting[0]                          //--> returns byte value of first char
	fmt.Println("First byte of greeting:", firstByte) // --> asic val of h
	// cobverintg byte ro string
	fmt.Println("First byte as string:", string(firstByte))

}
