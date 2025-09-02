package main

import (
	"fmt"
)

type Student struct {
	name   string
	age    int
	gender string
	grade  string
}

func (s Student) checkGrade() string {
	if s.grade == "A" || s.grade == "B" {
		return "pass"
	} else {
		return "fail"
	}
}

func main() {
	var students []Student

	students = append(students, Student{"John Doe", 20, "Male", "A"})
	students = append(students, Student{"Mochi", 12, "Female", "C"})
	students = append(students, Student{"Tiss", 13, "Male", "C"})

	fmt.Println("Welcome to student cli!")

	for i := 0; i < len(students); i++ {
		s := students[i]
		fmt.Println("Student Name: ", s.name)
		fmt.Println("Age: ", s.age)
		fmt.Println("Gender: ", s.gender)
		fmt.Println("Grade: ", s.grade)
		fmt.Println("Grade Check: ", s.checkGrade())
		if s.age >= 18 {
			fmt.Println("Student is an adult")
		} else {
			fmt.Println("Minor Student")
		}
	}
}
