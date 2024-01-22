package main

import (
	"fmt"
	"kelas-golang-pzn/go-dasar/database"
	_ "kelas-golang-pzn/go-dasar/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
