package main

import "fmt"

// interface kosong/any merupakan tipe data sembarang
func Ups() any {
	//return 1
	//return false
	return "upss"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}
