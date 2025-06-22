package main

import "fmt"

// Define the Car struct
type Car struct {
    Name  string
    Type  string
    Brand string
    Years int
}


//methods in structs
//-->  val reciver
func (c Car) DisplayDetails() {
    fmt.Printf("Car: %s, Brand: %s, Type: %s\n", c.Name, c.Brand, c.Type)
}
//c Car) is the receiver that binds the method DisplayDetails to the Car struct.
// The method can be called using the dot operator.

//--> pointer reciver
func (c *Car) UpdateName(newName string) {
    c.Name = newName
}	

// Method to check operational status
func (c Car) OperationalStatus() {
    if c.Years > 15 {
        fmt.Println(c.Name, "is no longer operational.")
    } else {
        fmt.Printf("%s has %d years of operation left.\n", c.Name, 15-c.Years)
    }
}
func main() {
    // Initialize a Car instance
    car := Car{
        Name:  "Model S",
        Type:  "Electric",
        Brand: "Tesla",
        Years: 3,
    }

    // Accessing fields
    fmt.Println("Car Name:", car.Name)
    fmt.Println("Years in Service:", car.Years)

      // Call methods
    car.DisplayDetails()
    car.OperationalStatus()

    // Update brand using pointer receiver
    car.UpdateName("SpaceX")
    fmt.Println("Updated Brand:", car.Brand)
}



