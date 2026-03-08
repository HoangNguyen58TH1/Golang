package timers_and_tickers

import (
	"fmt"
	"time"
)

func Timers() {
	// timers
	now := time.Now()
	fmt.Println("now:", now)

	// duration
	fmt.Println("duration:", 5*time.Nanosecond) // GO auto convet 5 --> time.Duration
	fmt.Println("duration:", 5*time.Microsecond)
	fmt.Println("duration:", 5*time.Millisecond)
	fmt.Println("duration:", 5*time.Second)
	fmt.Println("duration:", 5*time.Minute)
	fmt.Println("duration:", 5*time.Hour)
	fmt.Println("Day:", time.Now().Day())         // day of the month
	fmt.Println("Weekday:", time.Now().Weekday()) // day of the month

	// Duration with variable
	seconds := 10
	// NOT work because Go NOT allow nhân 2 kiểu khác nhau (int × time.Duration)
	// fmt.Println(seconds * time.Second)
	// variable int --> phải convert sang time.Duration
	fmt.Println(time.Duration(seconds)) // default nano second
	fmt.Println(time.Duration(seconds) * time.Second)
}

func Tickers() {
	// 1. print time after 1 second
	// ticker is a pointer to time.Ticker = *time.Ticker
	ticker1 := time.NewTicker(1 * time.Second)
	for t := range ticker1.C {
		fmt.Println("Tick 1 at: ", t)
	}

	// 2. stop ticker after 3 times
	ticker2 := time.NewTicker(1 * time.Second)
	count := 0

	for t := range ticker2.C {
		fmt.Println("Tick 2 at: ", t)
		count += 1 // count++
		if count == 3 {
			break
		}
	}

	// 3. Ticker with select:
	ticker3 := time.NewTicker(2 * time.Second) // ticker3 = *time.Ticker
	done := time.After(7 * time.Second)        // done = <-chan time.Time
	for {
		select { // listen 2 channels cùng lúc
		case t := <-ticker3.C: // (type: <-chan time.Time)
			fmt.Println("Tick at:", t)
		case <-done:
			fmt.Println("Stop program")
			ticker3.Stop() // ticker3 don't send tick to channel
			return         // exit function (otherwise deadlock!)
		}
	}
	// ┌────────────────────┐
	// │   main goroutine   │
	// │                    │
	// │   select {         │
	// │   <-ticker3.C      │
	// │   <-done           │
	// │   }                │
	// └─────────┬──────────┘
	// 					│
	// ┌─────────┴─────────────────────┐
	// │                               │
	// ┌─────────────────────┐       ┌─────────────────────┐
	// │ ticker goroutine    │       │ timer goroutine     │
	// │                     │       │                     │
	// │ every 2s            │       │ sleep 7s            │
	// │ send -> ticker3.C   │       │ send -> done        │
	// │                     │       │ exit                │
	// └─────────────────────┘       └─────────────────────┘
}
