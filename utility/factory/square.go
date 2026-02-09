package factory

import "fmt"

// concrete struct - init object dc
type Square struct{}

func (Square) Draw() {
	fmt.Println("--- Start draw ---")
	fmt.Println("Drawing a Square")
	fmt.Println("--- Done draw ---")
}
