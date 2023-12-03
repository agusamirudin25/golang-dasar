package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// implementasi interface dengan struct
type Person struct {
	Name string
}

func (person Person) GetName() string {
	// TODO implement me
	return person.Name
}

func main() {
	person := Person{Name: "Agus"}
	SayHello(person)
}
