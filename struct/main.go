package main

import "fmt"

func main() {
	type Car struct {
    Name           string
    Brand          string
    Type           string
    YearsInService int
}

//name field initalization
car1 := Car{
    Name:           "Model S",
    Brand:          "Tesla",
    Type:           "Electric",
    YearsInService: 3,
}
fmt.Println(car1.Name)

//Anonymous Structs --> one time use
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  30,
}

fmt.Println(person.Name) 


//embedding Struct
type Engine struct {
    Horsepower int
    Type       string
}

type CarWithEngine struct {
	Name    string
	Brand   string
	Type    string
	Engine 
	YearsInService int

}
 car2 := CarWithEngine{
        Name:           "Model S",
        Brand:          "Tesla",
        Type:           "Electric",
        YearsInService: 3,
        Engine: Engine{
            Horsepower: 670,
            Type:       "Dual Motor",
        },
    }

    fmt.Println(car2.Engine.Horsepower)

	//Comparing Structs

	  car4 := Car{"Model S", "Tesla", "Electric", 3}
    car5 := Car{"Model S", "Tesla", "Electric", 4}

    if car4 == car5 {
        fmt.Println("The cars are identical.")
    } else {
        fmt.Println("The cars are different.")
    }
}