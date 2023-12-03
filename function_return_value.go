package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Agus")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// cara menghiraukan return value
	_, lastNames := getFullName()
	fmt.Println(lastNames)
}

// return multiple values
func getFullName() (string, string) {
	return "Agus", "Amirudin"
}
