package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}

// var declares 1 or more variables.

// You can declare multiple variables at once.

// The := syntax is shorthand for declaring and initializing a variable.
// It is shorthand for var f string = "apple".
// The type of the variable will be inferred from the value on the right-hand side.
// Variables declared without a corresponding initialization are zero-valued.
// The zero value is a default value for a variable that has not been initialized.



// const declares a constant value.
// A const statement can appear anywhere a var statement can.

// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax.

//example of constants
const Pi = 3.14
const World = "world"
 const Truth = true
 func main() {
	 const Truth = true
	 fmt.Println("Hello", World)
	 fmt.Println("Happy", Pi, "Day")
	 fmt.Println("Go rules?", Truth)
 }
 // The const keyword is used to declare constants.


