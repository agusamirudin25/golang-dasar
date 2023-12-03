package helper

// access modifier pada golang cukup berbeda dengan bahasa pemrogarman pada umumnya
// di php misal, ada public private and protected
// kalau di golang, ditentukan dengan penamaan suatu variable atau function
// jika awalnya huruf kecil maka dia private (tidak bisa diakses oleh package lain)
// jika awalya huruf besar maka dia public (bisa diakses oleh package lain)

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
