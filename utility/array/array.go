package array

import "fmt"

func StudyArray() {
	// simple data type
	a := 10 // var a int = 10
	fmt.Println(a)

	// arrays
	var arr1 [3]int
	arr1[0], arr1[1] = 1, 2
	fmt.Println(arr1)
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// slices = array but can change length
	slice1 := []int{1, 2, 3, 4, 5}
	slice1 = append(slice1, 6)
	fmt.Println(slice1)
}
