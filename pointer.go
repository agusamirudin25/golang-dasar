package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Majalengka", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Bandung"

	fmt.Println("address1", address1)
	fmt.Println("address2", address2)
}
