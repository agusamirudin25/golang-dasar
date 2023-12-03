package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

func main() {
	sayHelloTo("Agus", "Amirudin")
}
