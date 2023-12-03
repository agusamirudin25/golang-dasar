package main

import "fmt"

/*
func defer digunakan untuk menjalankan fungsi setelah fungsi lain selesai dieksekusi
func panic digunakan untuk menghentikan aplikasi
func recover digunakan untuk menangkap message panic
*/

func endApplication() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi kesalahan", message)
}

func runApp(err bool) {
	defer endApplication()
	if err {
		panic("Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Agus")
}
