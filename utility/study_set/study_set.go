package study_set

import "fmt"

// type Set123 map[string]struct{} --> name Set define kieu nao cung dc
type Set map[string]struct{}

func StudySet() {
	s := make(Set)
	fmt.Println("set:", s)

	// add new item
	s["item1"] = struct{}{}
	s["item2"] = struct{}{}
	fmt.Println("set:", s)

	// check value
	_, ok := s["item2"]
	fmt.Println("ok:", ok)

	// delete item
	delete(s, "item2")
	fmt.Println("set:", s)

	// call func
	fmt.Println("getSetValues(s):", getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}

// quick implement like this
// set := make(map[string]struct{})

// Set = map[T]struct{}
// T = type = string, int, int64, bool, pointer
