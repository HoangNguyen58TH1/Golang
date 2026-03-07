package locks_and_syncs

import (
	"fmt"
	"sync"
	"time"
)

// 0. KO use Mutex to Lock ==> race condition --> nhiều Goroutine update trùng nhau do KO LOCK
func UnLock() {
	var counter int
	for i := 0; i < 1000; i++ {
		go func() {
			counter++ // read, +1, write
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter) // result random, NOT equal 1000
}

// 1. sync.Mutex
func LockWithMutex() {
	var counter int
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(counter) // always = 1000
}

// 2 + 3. sync.RWMutex – nhiều reader, 1 writer
var data = make(map[string]string)
var mu sync.RWMutex
var wg sync.WaitGroup

func LockWithRWMutex() {
	data["name"] = "Hoang"
	for i := 1; i <= 3; i++ { // 3 reader
		wg.Add(1)
		go read(i, "name")
	}
	wg.Add(1)
	go write("name", "Toni")  // 1 writer
	for i := 4; i <= 5; i++ { // add more reader after writer
		wg.Add(1)
		go read(i, "name")
	}
	wg.Wait()
}

func read(id int, key string) {
	defer wg.Done()
	mu.RLock()
	fmt.Println("Reader", id, "reading:", data[key])
	time.Sleep(200 * time.Millisecond) // giả lập đọc chậm
	mu.RUnlock()
}

func write(key, value string) {
	defer wg.Done()
	mu.Lock()
	fmt.Println("Writer writing:", value)
	data[key] = value
	time.Sleep(500 * time.Millisecond) // giả lập ghi chậm
	mu.Unlock()
}

// Reader 5 reading: Hoang
// Reader 3 reading: Hoang
// Writer writing: Toni
// Reader 4 reading: Toni
// Reader 1 reading: Toni
// Reader 2 reading: Toni

// 4. sync.Once
func SyncOnce() {
	var once sync.Once
	for i := 0; i < 5; i++ {
		go once.Do(writeLogger) // Dù có 5 goroutines → Init DB only 1
	}
	time.Sleep(time.Second)
}

func writeLogger() {
	fmt.Println("== method SyncOnce run completed ...! ==")
}
