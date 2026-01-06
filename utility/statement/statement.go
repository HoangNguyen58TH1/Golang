package statement

import "fmt"

func StudyStatement() {
	i := 5

	// IF ELSE
	if i < 0 {
		fmt.Println("negative number")
	} else if i == 0 {
		fmt.Println("zero number")
	} else {
		fmt.Println("positive number")
	}
	// SWITCH
	switch i := -1; {
	case i < 0:
		fmt.Println("negative number")
		fallthrough // only use in switch, NOT recheck condition of below case
	case i == 0:
		fmt.Println("zero number")
	default:
		fmt.Println("positive number")
	}

	fmt.Println("=====================")
	// switch os := runtime.GOOS; os {
	switch os := "windows"; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// windows,..
		fmt.Printf("OS is: %s\n", os)
		// fmt.Println(os)

		defer fmt.Println("Exiting function...1")
		defer fmt.Println("Exiting function...2")
		defer fmt.Println("Exiting function...3")
		fmt.Println("Entering function")
	}
}
