package main

import "fmt"

func main() {
	var aksi, konsekuensi1, konsekuensi2 string
	fmt.Scan(&aksi)
	if aksi == "mobil" {
		konsekuensi1 = "telat masuk kuliah"
	}
	if aksi == "motor" {
		konsekuensi1 = "bajunya akan basah"
	}
	if konsekuensi1 == "telat masuk kuliah" {
		konsekuensi2 = "mendapatkan pengurangan nilai"
	}
	if konsekuensi1 == "bajunya akan basah" {
		konsekuensi2 = "harus membawa baju salin"
	}
	fmt.Println(konsekuensi1)
	fmt.Println(konsekuensi2)
}