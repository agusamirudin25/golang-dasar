package main

import "fmt"

func main() {
	name := "Amir"

	if name == "Agus" {
		fmt.Println("Halo Agus")
	} else if name == "Amir" {
		fmt.Println("Halo Amir")
	} else {
		fmt.Println("Boleh Kenalan")
	}

	// shorthand if statement
	if length := len(name); length > 1 {
		fmt.Println("Nama pendek")
	} else {
		fmt.Println("Nama Panjang banget")
	}

	// SWITCH

	switch name {
	case "Agus":
		fmt.Println("Halooo Agus")
	default:
		fmt.Println("Hmmm tidak ada")
	}

	// switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Aman")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
