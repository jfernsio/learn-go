package main

import (
	"fmt"
	"log"
	"os"
)

func readFile() {
	file, err := os.Open("temp.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	os.WriteFile("temp.txt", []byte("Hello, World!"), 0644)
	// Perform operations on the file
	log.Println("File opened successfully")
}

func main() {
	readFile()
	defer fmt.Println("Executed at the end of main")
	fmt.Println("Hello, Go!")

	defer fmt.Println("First deferred statement")
	defer fmt.Println("Second deferred statement")
	defer fmt.Println("Third deferred statement")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

}
