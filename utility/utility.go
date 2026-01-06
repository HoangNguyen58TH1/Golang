package utility

import "fmt"

func SayHello() {
	fmt.Println("Hello!")
}

func SayToni() { fmt.Println("Hello Toni!") }

//	func AddSubtractMultiple(a, b int) (int, int, int) {
//		return a + b, a - b, a * b
//	}
func AddSubtractMultiple(a, b int) (addition, subtraction, multiplication int) {
	addition = a + b
	subtraction = a - b
	multiplication = a * b
	return
}

func Sum(a, b int) int {
	return a + b
}

func ComputeMultiplyVal(val, a, b int, fn func(a, b int) int) int {
	return val * fn(a, b)
}
