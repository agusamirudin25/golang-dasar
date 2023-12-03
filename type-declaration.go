package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAgus NoKTP = "3210122312321"
	fmt.Println(ktpAgus)
	fmt.Println(NoKTP("32121212"))
}
