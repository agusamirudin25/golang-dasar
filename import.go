package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	result := helper.SayHello("Agus")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//fmt.Println(helper.version) // tidak bisa
	//fmt.Println(helper.sayGoodBye()) // tidak bisa
}
