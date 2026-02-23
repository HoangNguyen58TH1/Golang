package builder

import "fmt"

// Product
type User struct {
	name     string
	age      int
	isActive bool
}

// Builder
type UserBuilder struct {
	user *User
}

// Constructor
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &User{},
	}
}

// Method chaining
func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.user.name = name
	return b
}
func (b *UserBuilder) SetAge(age int) *UserBuilder {
	b.user.age = age
	return b
}
func (b *UserBuilder) SetActive(active bool) *UserBuilder {
	b.user.isActive = active
	return b
}

// Build
func (b *UserBuilder) Build() *User {
	return b.user
}

// step Product, Builder, Constructor, Method chaining, Build.
func BuilderPattern() {
	user := NewUserBuilder().
		SetName("Hoang").
		SetAge(30).
		SetActive(true).
		Build()

	fmt.Println(user)  // &{Hoang 30 true} --> return pointer, but display &value (& to know this is a poiter)
	fmt.Println(&user) // 0xc000012028 --> return memory address
	fmt.Println(*user) // {Hoang 30 true} --> return value
}

// var a *int --> declare a pointer
// &a ==> get address of pointer
// *a ==> get value of pointer
