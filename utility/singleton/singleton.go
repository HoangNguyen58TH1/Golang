package singleton

import (
	"fmt"
	"sync"
)

type Logger struct{}

func (*Logger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}

// biến global (package-level) --> Tồn tại suốt lifecycle app + use chung everywhere
var (
	instance *Logger
	once     sync.Once
	// Chỉ 1 goroutine được chạy block, dù có bao nhiêu goroutine, Các goroutine khác chờ hoặc skip
	// Không cần lock / mutex thủ công
	// đảm bảo 1 block code only run 1 time (thread-safe)
)

func GetLogger() *Logger {
	// once - check đã run chưa, chưa thì run, rồi thì skip
	once.Do(func() {
		fmt.Println("Create Logger instance")
		fmt.Println("instance:", instance)   // <nil>
		fmt.Println("&instance:", &instance) // 0x1142da8
		// fmt.Println("*instance:", *instance) // panic: runtime error: invalid memory address or nil pointer dereference
		instance = &Logger{}                 // tạo object + get memory address
		fmt.Println("instance:", instance)   // &{}
		fmt.Println("&instance:", &instance) // 0x1142da8
		fmt.Println("*instance:", *instance) // {}
	})
	return instance
}

func SinglePattern() {
	l1 := GetLogger()
	l2 := GetLogger()

	l1.Log("Hello")
	l2.Log("World")

	// cùng return về 1 instance
	fmt.Println(l1 == l2) // true == &{}
	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)

	fmt.Println(&l1 == &l2)  // false
	fmt.Println("&l1:", &l1) // 0xc0000a6018
	fmt.Println("&l2:", &l2) // 0xc0000a6028
	fmt.Println(*l1 == *l2)  // true
	fmt.Println("*l1:", *l1)
	fmt.Println("*l2:", *l2)
}

// instance → giữ object Singleton
// sync.Once → đảm bảo tạo 1 lần duy nhất
// GetLogger() → global access point

// sync.Once dùng cơ chế gì?
// Atomic + Mutex (lock)
// sync.Once = atomic check nhanh + lock mutex để đảm bảo chỉ 1 goroutine chạy

// Program start
//    ↓
// instance = nil
// once = chưa chạy
//    ↓
// GetLogger() lần 1
//    → once.Do() chạy
//    → instance được tạo
//    ↓
// GetLogger() lần 2
//    → once.Do() bị skip
//    ↓
// Luôn trả về cùng instance

// var a *int --> declare a pointer
// &a ==> get address of pointer
// *a ==> get value of pointer
