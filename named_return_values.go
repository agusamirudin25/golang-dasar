package main

import "fmt"

func getCompleteNames() (firstName, middleName, lastName string) {
	firstName = "Agus"
	middleName = "Amir"
	lastName = "Udin"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastname := getCompleteNames()
	fmt.Println(firstName, middleName, lastname)
}
