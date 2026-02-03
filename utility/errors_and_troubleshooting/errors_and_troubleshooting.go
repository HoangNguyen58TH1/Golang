package errors_and_troubleshooting

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var scMapping = map[string]int{
	"Hoang":   1,
	"Toni":    2,
	"VanCute": 3,
}

var ErrCrewNotFound = errors.New("crew member not found")

// define error struct with Error() func
type findError struct {
	Name, Server, Msg string
}

func (e findError) Error() string {
	return e.Msg
}

func findSC(name, server string) (int, error) {
	if v, ok := scMapping[name]; !ok {
		// return -1, errors.New("crew member not found")
		// return -1, ErrCrewNotFound
		// return -1, fmt.Errorf("crew member: %s could not be found on server: %s", name, server)
		// return -1, findError{name, server, "Crew member not found"}

		// raise exception (such as raise error in Rails)
		panic("Crew member not found")
	} else {
		return v, nil
	}
}

func MainFunction() {
	defer func() {
		fmt.Println("Run defer func.")
		if err := recover(); err != nil {
			fmt.Println("A panic recovered:", err)
		}
	}()

	name := "Hoang"
	// name := "invalid name"
	clearance, err := findSC(name, "server 1")

	if err != nil {
		fmt.Println("Error occured:", err)
		if v, ok := err.(findError); ok {
			fmt.Println("Name is:", v.Name)
			fmt.Println("Server name is:", v.Server)
		}
	} else {
		fmt.Println("Clearance level found:", clearance, ". Error code:", err)
	}
}

func StandardInput() {
	var a int
	var b string
	fmt.Fscan(os.Stdin, &a, &b) // enter a and b
	fmt.Println(a, b)
}

func ReadFile() {
	file, err := os.Open("utility/errors_and_troubleshooting/text.txt")
	fmt.Println("file:", file)

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	defer file.Close()

	var x int
	var y float64
	// read from file and get the first 2 tokens with type: int & float64
	// Fscan store tokens (separate white space, new line, a tab)
	// if can NOT parse token --> panic: expected integer
	_, err = fmt.Fscan(file, &x, &y)
	if err != nil {
		panic(err)
	}
	fmt.Println("x:", x, ", y:", y)
}

func CheckContains() {
	fmt.Println(strings.Contains("hoang", "hoa"))  // t
	fmt.Println(strings.Contains("hoang", "toni")) // f
	fmt.Println(strings.Contains("hoang", ""))     // t
	fmt.Println(strings.Contains("", ""))          // t

	// compare 2 strings in case-insensitive
	fmt.Println(strings.EqualFold("Go", "go"))  // t
	fmt.Println(strings.EqualFold("Goa", "go")) // f

	// Ruby: array.join(sep) --> split(' ') separate theo ' '
	// GO: strings.Join(array, sep)
	s := []string{"hoang", "toni", "van"}
	fmt.Println(strings.Join(s, ", ")) // hoang, toni, van
}
