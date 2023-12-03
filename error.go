package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	}

	return nilai / pembagi, nil
}

func main() {
	result, err := Pembagi(100, 0)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Hasil pembagian", result)
	}

}
