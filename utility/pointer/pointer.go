package pointer

import "fmt"

func StudyPointer() {
	// pointer = memory address of a varible
	// x = 5
	// &x --> get memory address (0xc000018168)
	// p1 *int --> declare a pointer (only receive and return memory address)
	// *p1 --> get original value

	var p1 *int      // declare a pointer
	i := 4           // create integer using type inference
	p1 = &i          // assign the address of i for the pointer
	fmt.Println(*p1) // 4
	fmt.Println(p1)  // 0xc000018168

	j := i + 1
	fmt.Println(&j) // 0xc000018168
	fmt.Println(j)  // 5

	i = 5
	fmt.Println(*p1) // 5
	fmt.Println(p1)  // 0xc000018168

	p2 := &i
	fmt.Println(p2)  // 0xc000018168
	fmt.Println(*p2) // 5

	// change origin value via pointer
	x := 10
	fmt.Println("-------------------")
	fmt.Println(x)
	changeX(&x)
	fmt.Println(x)
	fmt.Println("-------------------")
}

func changeX(x *int) {
	fmt.Println("-------")
	fmt.Println(x)
	fmt.Println(*x)
	*x = 11
	fmt.Println(x)
	fmt.Println(*x)
	fmt.Println("-------")
}
