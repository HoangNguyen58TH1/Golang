package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func StudyGoroutines() {
	// fmt.Println("Application start")
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println("Goroutines: ", i)
	// 	}
	// }()
	// fmt.Println("Application end")
	// time.Sleep(time.Second) --> remove this line --> goroutine NOT work

	var wg sync.WaitGroup
	fmt.Println("Application start")
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutines: ", i)
		}
		wg.Done()
	}()
	fmt.Println("Application end")
	wg.Wait()
}

// time.Sleep(time.Second) --> sleep 1s
// var wg sync.WaitGroup --> init sync.WaitGroup
// wg.Add(1) --> add 1 Goroutine
// wg.Done() --> make Done
// wg.Wait() --> program wait Goroutine run done rồi mới end program

// Channel --> là 1 đường ống connect các goroutines để share data vs nhau
// make(chan <type>)
// channelName <- ==> send data
// <- channelName ==> receive data
// Default send/receive giữa các goroutines bị BLOCK, until cả 2 đã ready để send/receive

func StudyGoroutinesChannel() {
	done := make(chan bool)
	fmt.Println("Application start")

	go func() {
		time.Sleep(time.Second)
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutines: ", i)
		}
		done <- true
	}()

	fmt.Println("Application end")
	<-done
}

func StudyGoroutinesChannelBuffering() {
	done := make(chan string, 1)
	fmt.Println("Application start")

	done <- "hoang"
	fmt.Println("Application end")
}

// send and receive channel such as param
func SendAndReceiveChannel() {
	channel := make(chan string, 64)
	go sendValue("Hello", channel)
	go sendValue("Xin chao", channel)

	go receiveValue(channel)
	time.Sleep(time.Second)
}
func sendValue(number string, channel chan<- string) {
	for {
		channel <- number
	}
}
func receiveValue(channel <-chan string) {
	for v := range channel {
		fmt.Println(v)
	}
}
