package main

import "fmt"

// menggunakan type declaration
type Filter func(string) string

func filterText(name string, filter Filter) string {
	return "Hello " + filter(name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	filter := spamFilter

	filtered := filterText("Anjing", filter)

	fmt.Println(filtered)

}
