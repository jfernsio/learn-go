package main

import ("fmt")

func  main () {
	//A rune is declared using single quotes (' '),
	//  and its value is the Unicode code point of the character:

	var ch rune = 'A'
fmt.Println(ch)         // Output: 65 (Unicode code point of 'A')
fmt.Println(string(ch)) // Output: A

//unicode exmpls
 s:= "Hello 🌍" 

 //length in bytes n runes
 fmt.Println("length in bytes : ",len(s))
  fmt.Println("length in runes : ",len([]rune(s)))

  for _, r := range []rune(s) {
        fmt.Printf("Rune: %c\n", r)
    }
}