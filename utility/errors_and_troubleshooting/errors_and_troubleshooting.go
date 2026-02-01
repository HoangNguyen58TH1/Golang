package errors_and_troubleshooting

import (
	"errors"
	"fmt"
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
