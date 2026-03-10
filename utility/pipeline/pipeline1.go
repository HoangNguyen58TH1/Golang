package pipeline

import (
	"fmt"
)

// chan int   → channel truyền int, có thể send + receive
// chan<- int → send-only channel (chỉ ghi)
// <-chan int → receive-only channel (chỉ đọc)

// stage 1: generate numbers
func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// stage 2: square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func PipelinePattern() {
	// ch1 := generator(1, 2, 3, 4)
	// ch2 := square(ch1)

	ch1 := generator(1, 2, 3, 4)
	ch2 := square(ch1)

	for result := range ch2 {
		fmt.Println(result)
	}
}

// generator goroutine
//     │
//     │ send
//     ▼
// channel ch1
//     │
//     ▼
// square goroutine
//     │
//     │ send
//     ▼
// channel ch2
//     │
//     ▼
// main goroutine (consumer)

// generator → produce data
// square → transform data
// main → consume data
