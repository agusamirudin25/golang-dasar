package main

import "fmt"

func main() {
	//var person = map[string]string{}
	//person["name"] = "Agus"
	//person["address"] = "Majalengka"

	var person = map[string]string{
		"name":    "Agus Amirudin",
		"address": "Majalengka",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
