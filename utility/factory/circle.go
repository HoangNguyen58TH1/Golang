package factory

import "fmt"

// concrete struct - init object dc
type Circle struct{}

// struct Circle sẽ có method Draw()
func (Circle) Draw() {
	fmt.Println("--- Start draw ---")
	fmt.Println("Drawing a Circle")
	fmt.Println("--- Done draw ---")
}
