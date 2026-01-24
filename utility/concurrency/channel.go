package concurrency

import (
	"fmt"
)

func UseChannelToCommunicateBetweenGoroutines() {
	// Unbuffered channel
	c := make(chan bool)
	go waitAndSay(c, "World")
	fmt.Println("Hello")
	c <- true // we send a signal to c in orderto allow waitAndSay to continue
	<-c       // we wait to receive another signal on c before we exit
}

func waitAndSay(c chan bool, s string) {
	if b := <-c; b { // if <statement>; <condition> {
		fmt.Println(s)
	}
	// b := <-c
	// if b {
	// 	fmt.Println(s)
	// }
	c <- true
}

// ch := make(chan <type>, <number of buffered channels>)
// Sender only block when the buffer is full
// Receivers block when the buffer is empty

func UseChannelWithBuffer() {
	// Buffered channel
	ch := make(chan string, 5) // make(chan <type>, <number of buffered channels>)

	ch <- "Hello"
	ch <- "Hoang"
	ch <- "Toni"

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	close(ch) // close the channel ch
	for value := range ch {
		fmt.Println(value)
	}

	v, ok := <-ch // ok is channel is open/close?
	fmt.Println("Channel open?", ok, ", value:", v)
	fmt.Println("Channel close?", !ok, ", value:", v)
}
