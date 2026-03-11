package laws_of_reflection

import (
	"fmt"
	"reflect"
)

// 1. variable → interface{} → reflect.Type / reflect.Value
func ValueToReflect() {
	x := "Hoang"
	fmt.Println(reflect.TypeOf(x))  // string
	fmt.Println(reflect.ValueOf(x)) // Hoang

	type User struct {
		Name string
		Age  int
	}
	u := User{"Alice", 30}
	v := reflect.ValueOf(u)
	fmt.Println("u:", u)                   // {Alice 30}
	fmt.Println("u:", reflect.TypeOf((u))) // laws_of_reflection.User (Struct)
	fmt.Println("v:", v)                   // {Alice 30}
	fmt.Println("v:", reflect.TypeOf((v))) // reflect.Value
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i)) // Alice 30
	}
	fmt.Println("========================")
}

func ReflectToValue() {
	x := 8 // x = orgin_value
	v := reflect.ValueOf(x)
	fmt.Println("v:", v)                                    // 8
	fmt.Println("reflect.TypeOf((v):", reflect.TypeOf((v))) // reflect.Value

	orgin_value := v.Interface()                                                // convert from reflection object --> value
	fmt.Println("orgin_value:", orgin_value)                                    // 8
	fmt.Println("reflect.TypeOf((orgin_value):", reflect.TypeOf((orgin_value))) // int
	fmt.Println("========================")
}

func SetReflectValue() {
	x := 10
	fmt.Println("x:", x)                                 // 10
	fmt.Println("reflect.TypeOf(x):", reflect.TypeOf(x)) // int

	fmt.Println("reflect.ValueOf(&x):", reflect.ValueOf(&x)) // 0xc0000180c0
	v := reflect.ValueOf(&x).Elem()
	fmt.Println("v:", v)                                 // 10
	fmt.Println("reflect.TypeOf(v):", reflect.TypeOf(v)) // reflect.Value

	v.SetInt(20)                                         // update value
	fmt.Println("x:", x)                                 // 20
	fmt.Println("reflect.TypeOf(x):", reflect.TypeOf(x)) // int
}
