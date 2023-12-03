package main

import "fmt"

func main() {
	var name [2]string

	name[0] = "Agus"
	name[1] = "Amirudin"

	fmt.Println(name[0])
	fmt.Println(name[1])

	// langsung inisialisasi data
	var values = [3]int{20, 30, 40}

	fmt.Println(values)

	// gunakan [...] untuk tidak menentukan panjang array diawal
	// [...]int{10, 20, 30}
}
