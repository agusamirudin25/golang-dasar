package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

// struct method = func yang menempel pada struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var agus Customer

	fmt.Println(agus)
	agus.Name = "Agus Amirudin"
	agus.Address = "Majalengka"
	agus.Age = 27

	fmt.Println(agus.Name, agus.Address, agus.Age)

	// struct literals
	joko := Customer{
		Name:    "Joko",
		Address: "Karawang",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 50}
	fmt.Println(budi)

	// cara manggil method struct
	budi.sayHello("Resti") // Hello Resti my name is Budi
	joko.sayHello("Resti") // Hello Resti my name is Joko
	agus.sayHello("Resti")
}
