package laws_of_reflection

import (
	"fmt"
	"reflect"
)

func ReflectWithStruct() {
	type User struct {
		Name string
		Age  int
	}
	user := User{"Alice", 30}

	fmt.Println("user:", user)                  // {Alice 30}
	fmt.Println("user:", reflect.ValueOf(user)) // {Alice 30}
	fmt.Println("user:", reflect.TypeOf(user))  // laws_of_reflection.User (Struct)
	fmt.Println("----------")

	// VALUE OF STRUCT
	userRValue := reflect.ValueOf(user)
	fmt.Println("userRValue:", userRValue)                  // {Alice 30}
	fmt.Println("userRValue:", reflect.TypeOf(userRValue))  // reflect.Value
	fmt.Println("userRValue:", reflect.ValueOf(userRValue)) // <laws_of_reflection.User Value>
	for i := 0; i < userRValue.NumField(); i++ {
		fmt.Println(userRValue.Field(i)) // Alice 30
	}
	fmt.Println("----------")

	// TYPE OF STRUCT
	userRType := reflect.TypeOf(user)
	fmt.Println("userRType:", userRType)                  // laws_of_reflection.User (Struct)
	fmt.Println("userRType:", reflect.TypeOf(userRType))  // *reflect.rtype
	fmt.Println("userRType:", reflect.ValueOf(userRType)) // laws_of_reflection.User
	for i := 0; i < userRType.NumField(); i++ {
		fmt.Println(userRType.Field(i))

		fmt.Println(userRType.Field(i).Name)
		fmt.Println(userRType.Field(i).Type)
		fmt.Println(userRType.Field(i).Offset)
		fmt.Println(userRType.Field(i).Index)
		fmt.Println(userRType.Field(i).Anonymous)
	}
}

// 									{Name   string  0    [0]  false}
//					 				{Age    int     16   [1]  false}
// StructField -> Name(string)  Type(Type)  Offset(uintptr)  Index([]int)  Anonymous(bool)
