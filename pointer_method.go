package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	agus := Man{"Agus"}
	agus.Married()

	fmt.Println(agus.Name)
}
