package file_handling

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func HandleJsonFile() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Job  string `json:"job"`
		Sex  string `json:"sex"`
	}

	file, err := os.ReadFile("utility/file_handling/user.json")
	if err != nil {
		fmt.Println("Can NOT read file utility/file_handling")
	}

	var user User
	err = json.Unmarshal(file, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Sex:", user.Sex)
	fmt.Println("Job:", user.Job)
	fmt.Println("=========================")
}

func HandleXMLFile() {
	type User struct {
		Name string `xml:"Name"`
		Age  int    `xml:"Age"`
		Job  string `xml:"Job"`
		Sex  string `xml:"Sex"`
	}

	file, err := os.ReadFile("utility/file_handling/user.xml")
	if err != nil {
		fmt.Println("Can NOT read file utility/file_handling")
	}

	var user User
	err = xml.Unmarshal(file, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Sex:", user.Sex)
	fmt.Println("Job:", user.Job)
	fmt.Println("=========================")
}

func HandleCSVFile() {
	file, _ := os.Open("utility/file_handling/user.csv")
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("err:", err)
	}

	for _, record := range records {
		fmt.Println(record[0], record[1], record[2], record[3])
	}
}
