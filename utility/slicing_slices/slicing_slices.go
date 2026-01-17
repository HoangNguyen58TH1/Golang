package slicing_slices

import "fmt"

func SlicingSlices() {
	// declare with memory + value
	s1 := []int{1, 2, 3}

	// declare with memory + no value
	s2 := []int{}

	// declare no memory + no value
	var s3 []int

	// declare memory + value + limit min/max capacity
	s4 := make([]int, 5, 10) // make([]T, length, capacity)

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	s4[0] = 1
	s4[1] = 2
	s4[2] = 3
	s4[3] = 4
	s4[4] = 5
	fmt.Printf("s4: %d, len(s4): %d, cap(s4):: %d \n", s4, len(s4), cap(s4)) // [], 5, 10

	// s4[5] = 6 --> error index out of range [5] --> Invalid
	s4 = append(s4, 6, 7, 8, 9, 10)                                          // Valid
	fmt.Printf("s4: %d, len(s4): %d, cap(s4):: %d \n", s4, len(s4), cap(s4)) // [], 10, 10

	s4 = append(s4, 11)                                                      // auto allocated thêm x2 capacity
	fmt.Printf("s4: %d, len(s4): %d, cap(s4):: %d \n", s4, len(s4), cap(s4)) // [], 11, 20

	s4 = s4[2:4]
	fmt.Printf("s4: %d, len(s4): %d, cap(s4):: %d \n", s4, len(s4), cap(s4)) // [3, 4], 2, 18
	// capacity = 18 because nó tính từ vị trí low(first) of slice.

	fmt.Println("What s4 actually see: ", s4[:cap(s4)]) // [3 4 5 6 7 8 9 10 11 0 0 0 0 0 0 0 0 0]
	// s4[:cap(s4)] = 18 --> s4[:18] --> from low to 18
}

// array (size 10): [0][0][0][0][0][_][_][_][_][_]
//                   ↑              ↑
//                 len=5          cap=10
// s4[5] = 6 --> nó đang tìm index 5 rồi nhét value=6 vào đó --> FAILED
// s4 = append(s4, 6) --> check capacity, update value, update length

// s4 := make([]int, 5, 10)
// s4[0] = 1
// s4[1] = 2
// s4[2] = 3
// s4[3] = 4
// s4[4] = 5
// s4 = append(s4, 6, 7, 8, 9, 10)
// Vậy thì khi s4 = append(s4, 11) tôi check length=10, capa=20
// It mean nó tự cấp phát thêm bộ nhớ capaticy là 10 nữa hả?

func PointerSlice() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("s1:", s1) // [1 2 3 4 5 6]
	s2 := s1[2:4]
	s2[0] = 10             // s2[0] = s1[2]
	fmt.Println("s1:", s1) // [1 2 10 4 5 6]
}

func CopyElement() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 2)
	n := copy(s2, s1[2:4])
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("value of n (length)", n)
	s2[0] = 10
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
}

func RemoveElement() {
	a := []int{1, 2, 3, 4, 5}
	// remove a item (index=1)
	a = append(a[:1], a[2:]...)
	fmt.Println("a:", a)

	b := []int{1, 2, 3, 4, 5}
	// remove items (index=1, index=2)
	b = append(b[:1], b[3:]...)
	fmt.Println("b:", b)
}
