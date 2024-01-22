package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusNilaiAbsensi bool = nilaiAbsensi > 80

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsensi
	fmt.Println(lulus)
}
