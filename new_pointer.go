package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.City = "Majalengka"

	fmt.Println("alamat1", alamat1)
	fmt.Println("alamat2", alamat2)
}
