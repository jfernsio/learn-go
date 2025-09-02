package main 

// Struct Receiver Type: The method operates on a copy of the struct.
// Changes made within the method do not affect the original struct.
// Pointer Receiver Type: The method operates on a reference to the original struct.
// Changes made within the method affect the original struct.

import "fmt"

//struct def
type Rect struct {
	Width int
	Length int
}

//method with struct receiver
func (r Rect) Area() int {
	return  r.Width * r.Length
}

//method with pointer receiver
func (r *Rect) Scale(factor int) {
	r.Width *= factor
	r.Length *= factor
}

func main() {
	rec := Rect{Width: 10, Length: 5}

	//call area method struct rec
	fmt.Println("Inital Area : ", rec.Area())

	//call scale method pointer rec
	rec.Scale(10)

	//updated areas
	fmt.Println("Scaled Area:", rec.Area())
    fmt.Println("Updated Width:", rec.Width)
    fmt.Println("Updated Length:", rec.Length)

}

