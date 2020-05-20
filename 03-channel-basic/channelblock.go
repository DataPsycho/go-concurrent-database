package main

import (
	"fmt"
)

func channelblock() {
	ch := make(chan int)
	// We always need a sender and recever
	// to handle message in to a channel
	// Here we are trying to receive a message but we have not generated one
	fmt.Println(<-ch)
	// This line will never run
	ch <- 42
}

func channelblock_2() {
	ch := make(chan int)
	// We always need a sender and reveber
	// to oass nesage in to a channel
	// Here we can not reveive a message untill there is a recever
	// In that point no one is listening from that channel so we can not
	// Put anything
	ch <- 42
	fmt.Println(<-ch)

}

// Thats why go routine comes handy because if channel does not listen
// The go routine will go to sleep and wait for the listener from the channel
// Because anoter task waiting for a result go run time can run the second go routine
// by overpassing the first one
