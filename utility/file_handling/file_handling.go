package file_handling

import (
	"fmt"
	"os"
	"time"
)

func FileHandling() {
	file, err := os.Create("utility/file_handling/new_file.txt")
	if err != nil {
		fmt.Println("Can NOT create file!")
	}

	n, err := file.WriteString("hoang toni")
	fmt.Println("n:", n)
	if err != nil {
		fmt.Println("Can NOT write file!")
	}

	fmt.Println("file:", file)
	fmt.Println("file.Name():", file.Name())
	data, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Can NOT read file!")
	}
	fmt.Println("string(data):", string(data))

	time.Sleep(3 * time.Second)
	os.Remove(file.Name())
	fmt.Println("=========================")
}
