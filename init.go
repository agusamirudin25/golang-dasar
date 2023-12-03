package main

// gunakan "_" untuk blank identifier, artinya hanya akan menjalankan func init saja di package tersebut
import (
	"fmt"
	"golang-dasar/database"
	_ "golang-dasar/internal"
)

func main() {
	fmt.Println(database.GetConnection())
}
