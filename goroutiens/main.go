package main

import (
	"fmt"
	"sync"
	"time"
)

// ..basic go routine
func printNumbers () {
	for i :=1;1<=10;i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main () {
	go printNumbers()
    fmt.Println("Started printing numbers")
    time.Sleep(11 * time.Second) // Wait for the goroutine to finish
    fmt.Println("Finished printing numbers")
}

//managing go routines
// ..using wait groups

func main() {
    var wg sync.WaitGroup
    wg.Add(1) // Increment the WaitGroup counter
    
    go func() {
        defer wg.Done() // Decrement the counter when the goroutine completes
        printNumbers()
    }()
    
    wg.Wait() // Wait for all goroutines to complete
    fmt.Println("All goroutines completed.")
}

func printNumbers() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
// The sync.WaitGroup is used to wait for the program to finish all goroutines explicitly before exiting.

