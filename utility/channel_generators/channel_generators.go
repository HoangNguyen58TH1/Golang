package channel_generators

import "fmt"

func generator() <-chan int { // producer
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func Consumer() {
	for v := range generator() {
		fmt.Println(v)
	}
}

// goroutine (produce data)
//         │
//         ▼
//      channel
//         │
//         ▼
// consumer (main / function khác)
