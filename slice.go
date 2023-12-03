package main

import "fmt"

func main() {
	// slice merupakan potongan dari array
	// dan bentuk data serta total nya bisa berubah (dinamis)
	names := [...]string{"Agus", "Amir", "Udin", "Joko", "Budi", "Nugraha"}

	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	slice2 := names[:3]
	fmt.Println(slice2) // [Joko Budi Nugraha]

	slice3 := names[3:]
	fmt.Println(slice3) // [Joko Budi Nugraha]

	slice4 := names[:]
	fmt.Println(slice4) // [Agus Amir Udin Joko Budi Nugraha]

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	// cek days
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println("daySlice2 =", daySlice2)

	// cek days lama
	fmt.Println("days =", days)

	// membuat slice dari awal
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Agus"
	newSlice[1] = "Agus"
	//newSlice[2] = "Agus" // error, harus menggunakan `append`

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println("iniArray", iniArray)
	fmt.Println("iniSlice", iniSlice)

}
