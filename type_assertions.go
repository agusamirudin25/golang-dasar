package main

import "fmt"

func random() any {
	return false
}

func main() {
	var result any = random()
	//fmt.Println(result)

	// konversi menjadi string
	//var resultString string = result.(string)
	//fmt.Println(resultString)

	// konversi menjadi int
	//var resultInt int = result.(int)
	//fmt.Println(resultInt) // panic error

	// cara handle menggunakan switch
	switch typeValue := result.(type) {
	case string:
		fmt.Println("String", typeValue)
	case int:
		fmt.Println("Int", typeValue)
	default:
		fmt.Println("Unknown", typeValue)
	}
}
