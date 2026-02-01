package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

var scMapping = map[string]int{
	"Hoang":   1,
	"Toni":    2,
	"VanCute": 3,
}

func findSC(name, server string, c chan int) {
	random_time := rand.Intn(10)
	fmt.Println("random_time:", random_time, "seconds, of:", server)
	time.Sleep(time.Duration(random_time) * time.Second) // sleep a time from 0 --> 9 seconds

	c <- scMapping[name] // return security clearance from map
}

func StudySelect() {
	// make sure random ko return trùng value. Ttrong GO version mới (1.20) don't need, nhưng nó clear ý đồ mình muốn làm
	// rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)

	name := "Hoang"
	go findSC(name, "server 1", ch1)
	go findSC(name, "server 2", ch2)

	select {
	case sc := <-ch1:
		fmt.Println(name, "Has a security clearance of:", sc, "found in server 1")
	case sc := <-ch2:
		fmt.Println(name, "Has a security clearance of:", sc, "found in server 2")
	case <-time.After(4 * time.Second):
		fmt.Println("Search time out...!")
	}
}

// Select statement definition
// - wait multiple channels at the same time
// - block until 1 channel is ready
// - if multiple channels are ready, select pick on at ramdom

// t := time.Now()
// fmt.Println(t)
// fmt.Println(t.Unix())
// fmt.Println(t.UnixNano())

// 1s --> 1K ms
// 1s --> 1M us (microsecond)
// 1s --> 1B ns (nanosecond)
