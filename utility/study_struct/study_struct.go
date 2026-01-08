package study_struct

import "fmt"

// declare a struct
func DeclareStruct() {
	// 1
	cm1 := CMember{"Toni", 28, "Nha Chang", "Senior", 113}
	fmt.Println("cm1: ", cm1)

	// 2
	cm2 := CMember{
		name:      "Toni",
		age:       29,
		address:   "Da Nang",
		rank:      "Middle",
		clearance: 114,
	}
	fmt.Println("cm2: ", cm2)

	// 3
	var cm3 CMember
	cm3.name = "Tonii"
	cm3.age = 30
	cm3.address = "HCM"
	// cm3.rank = "Tech Lead"
	// cm3.clearance = 115
	fmt.Println("cm3: ", cm3)

	// Pointer
	pcm3 := &cm3
	pcm3.age = 31
	cm3.age = 32
	fmt.Println("cm3: ", cm3)
	// attrs
	fmt.Println("cm3.name: ", cm3.name)
	fmt.Println("&cm3.name: ", &cm3.name)
	fmt.Println("pcm3: ", pcm3)

	// Slices
	var crew []CMember
	crew = append(crew, cm1, cm2, cm3, CMember{"Van", 18, "DK", "senior", 9})
	fmt.Println("crew: ", crew)

	// for loop --> range
	for index, value := range crew {
		fmt.Println(index, value)
	}
}

// Slices
func StudySlices() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s := []int{2, 3, 5, 7, 11, 13}
	s1 := s[1:4]            // ==> [...)
	fmt.Println("s1: ", s1) // 1 2 3

	s2 := s[:2]             // ==> [...)
	fmt.Println("s2: ", s2) // 0 1

	s3 := s[1:]             // ==> [...)
	fmt.Println("s3: ", s3) // 1 2 3 4 5 6 7 8 9

	a := make([]int, 10)
	fmt.Println(a) // [0 0 0 0 0 0 0 0 0 0]

	for i, v := range a {
		fmt.Println(i, v)
	}
	for _, v := range a {
		fmt.Println("value: ", v)
	}
}

// Maps
func StudyMap() {
	// var m map[string]int --> KO có memory, nil --> lúc này gán key = value là error
	// m := make(map[string]int) --> đã có memory là {} rỗng --> gắn key=value work
	// m := map[string]int{} --> đã có memory là {} rỗng --> gắn key=value work

	// map() --> cấp phát hash table
	// make() --> cấp phát bộ nhớ cho map, cấp phát bộ nhớ xong mới đc gắn key = value
	// make() dùng để khởi tạo các kiểu reference type --> map, slice, channel

	cm := CMember{"Toni", 28, "Nha Chang", "Senior", 113}

	// var m map[string]CMember AND m = make(map[string]CMember)

	// use Literal
	// m := map[string]CMember{
	// 	"a": cm,
	// 	"b": cm,
	// }

	// Use make()
	m := make(map[string]CMember)
	m["Toni"] = cm
	m["Hon"] = cm
	fmt.Println("m: ", m)

	// Pointer
	n := m
	n["Van"] = CMember{"Van", 18, "DK", "Senior", 10}
	fmt.Println("n: ", n) // == fmt.Println("m: ", m)
	fmt.Println("m: ", m)

	// retrive
	k := m["Van"]
	fmt.Println("k: ", k) // {Van 18 DK Senior 10} --> value

	// check if the value exists?
	v, ok := m["Van"]
	fmt.Println("v: ", v)   // {Van 18 DK Senior 10} || { 0   0}
	fmt.Println("ok: ", ok) // true || false

	// delete key:value
	fmt.Println("len(m):", len(m))
	delete(m, "Hon")
	fmt.Println("m: ", m)
	fmt.Println("len(m):", len(m))

	// for loop Map
	for k, v := range m {
		fmt.Println("key:", k, "AND", "value:", v)
		fmt.Println("==========")
		v.PrintSecurityClearance()
		v.PrintNameAndAge()
		fmt.Println("==========")
	}
}

type CMember struct {
	name      string
	age       int
	address   string
	rank      string
	clearance int
}

// This is a METHOD of CMember, dùng value receiver(cm) to read AND in data
func (cm CMember) PrintSecurityClearance() {
	fmt.Println("cm.clearance:", cm.clearance)
}
func (cm CMember) PrintNameAndAge() {
	fmt.Println("Name is:", cm.name, "AND Age is:", cm.age)
}

// cm := CMember{"Toni", 28, "Nha Chang", "Senior", 113}
// fmt.Println("===")
// fmt.Println("call k.PrintSecurityClearance:", cm.PrintSecurityClearance())
// fmt.Println("===")
