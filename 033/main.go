// Exercise: Channels

// It's important to get this concept, document yourself first! :) 
// There are two concurrent routines (f1 and f2)
// Send a message "Hello from f1" from f1 function
// Receive the message from f1 into the f2 function, and print "I am f2 and ..." + the message from f1
// This should be done with a channel

package main

import "fmt"
import "time"

func f1 (c chan string){
	c <- "Hello from f1"
}

func f2 (c chan string){
	word := <-c
	fmt.Printf("I am f2 and %s\n", word)
}
func main () {
	c1 := make(chan string)
	go f1(c1)
	go f2(c1)

	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	time.Sleep(1 * time.Second)
}