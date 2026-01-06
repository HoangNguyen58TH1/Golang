package study_float

import (
	"fmt"
	"reflect"
)

func StudyFloat() {
	n := 32
	f := float64(n)
	fmt.Println(n)
	fmt.Println("reflect.TypeOf(n): ", reflect.TypeOf(n)) // int
	fmt.Println(f)
	fmt.Println("reflect.TypeOf(f): ", reflect.TypeOf(f)) // float64
	const pi = 3.14
	fmt.Println(pi)
	fmt.Println("reflect.TypeOf(pi): ", reflect.TypeOf(pi)) // float 64
	fmt.Println(pi * f)
	fmt.Println("reflect.TypeOf(pi*f): ", reflect.TypeOf(pi*f)) // float 64
}
