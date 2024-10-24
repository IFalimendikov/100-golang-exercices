// Exercise: Channels synchronisation

// We have a goroutine task, learn how the channel synchronisation works and make the program wait for the asynchronous function (make it synchronous ;) )
// This exercise is easy, you just have to add one single line

package main

import "fmt"
import "time"

func task(done chan bool) {
    fmt.Print("running...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main () {
	var done chan bool = make(chan bool)	
	go task(done)
    <- done

	// Your code goes here
	
}