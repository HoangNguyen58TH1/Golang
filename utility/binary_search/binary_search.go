package binary_search

import "fmt"

func IterationBinarySearch(array []int, target int) int {
	l := 0
	r := len(array) - 1

	for l <= r {
		mid := l + (r-l)/2

		if array[mid] == target {
			fmt.Printf("Found target in array with value is: %d at index: %d\n", array[mid], mid)
			return mid
		} else if array[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	fmt.Println("Can NOT found target in array!")
	return -1
}

func RecursiveBinarySearch(array []int, target, lowIndex, highIndex int) int {
	if highIndex < lowIndex {
		fmt.Println("Can NOT found target in array!")
		return -1
	}

	// mid := (lowIndex + highIndex) / 2
	mid := lowIndex + (highIndex-lowIndex)/2
	if array[mid] == target {
		fmt.Printf("Found target in array with value is: %d at index: %d\n", array[mid], mid)
		return mid
	} else if array[mid] < target {
		return RecursiveBinarySearch(array, target, mid+1, highIndex)
	} else {
		return RecursiveBinarySearch(array, target, lowIndex, mid-1)
	}
}
