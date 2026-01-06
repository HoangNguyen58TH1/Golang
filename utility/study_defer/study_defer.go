package study_defer

import "fmt"

func StudyDefer() {
	n := 1
	fmt.Println("===start===")
	for n <= 50 {
		n += n // n = n*2
		defer fmt.Println("n is: ", n)
	}
	fmt.Println("Result: ", n) // 64
	fmt.Println("===done===")
}
