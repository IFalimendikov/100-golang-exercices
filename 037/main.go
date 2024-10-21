// Exercise: Channels directions (only read/rx)

// Make a goroutine with a channel for only receive data.
// The function should be called "receive" and the receive-only channel should be it's 1st and only argument
// Sending data from that channel is prohibited / will cause compiler errors
// Feed some string into that channel.


package main

import "fmt"

func receive(ch <- chan string) {
	str := <- ch
	fmt.Print(str)
}

func main () {
	c1 := make(chan string, 1)

	c1 <- "Hey"
	receive(c1)
}