package main

import (
	"fmt"
	"sync"
	"time"
)

//wait grps used to sync go routines
//providede by sync package

//methods

// 1. add(n int)   ---> increment counter
// 2. done()       ---> decrement counter used w defer keyword
// 3. wait()


func worker (i int , wg *sync.WaitGroup) {
	defer 	wg.Done()
	fmt.Printf("Worker No  %v starting\n", i)
	
	time.Sleep(2 *time.Second)
	fmt.Printf("Worker NO %v completed\n",i)
}

func main () {
var wg  sync.WaitGroup

	for i := 0;i<3 ; i++ {
		wg.Add(1)
		go worker(i+1, &wg)
	}

	wg.Wait()
	fmt.Println("main process compelted")

}