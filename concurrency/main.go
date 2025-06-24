package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("gello from go routine")
}

func main () {
	go sayHello() //=> strt a go routine
	time.Sleep(2 * time.Second) // allow go routine to complete
	fmt.Println("Main func executed")
}

//Channels s provide a mechanism for goroutines to communicate and share data safely.
//  They ensure synchronized data transfer, maintaining the integrity of shared information.

//Unbuffered  -->requires the sender and receiver to be available simultaneously.

func sendData(channel chan int) {
    channel <- 42 // Send data into the channel
}

func main () {
	channel := make(chan int)
	go sendData(channel)
	fmt.Println(<-channel)
}

//Buffered --> allows data to be sent even if the receiver is not available.

func main() {
    channel := make(chan int, 3) // Create a buffered channel with capacity 3

    channel <- 10
    channel <- 20
    channel <- 30

    fmt.Println(<-channel) 
    fmt.Println(<-channel) 
    fmt.Println(<-channel) 
}