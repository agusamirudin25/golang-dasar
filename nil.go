package main

import "fmt"

/*
nil merupakan variable yang datanya kosong;
nil hanya bisa digunakan untuk tipe data interface, function, map, slice, pinter
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			name: name,
		}
	}
}

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
