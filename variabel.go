package main

import "fmt"

func main() {
	var hello string = "Ahmad Faqih"

	fmt.Println(hello)
	fmt.Println(hello)

	hello = "Faqih Nazafarin"

	fmt.Println(hello)

//	membuat variabel tidak menampilkan pertama

	var nama string
	nama = "Risma melani"

	fmt.Println(nama)

//	shorcut variabel

	namaLengkap := "Khumaira Nazafarin"
	fmt.Println(namaLengkap)

	nilai := 10
	fmt.Println(nilai)

//	Variabel constants = variabel yang tidak bisa dirubah nilanya

	const perusahaan string = "Workin.id"
	fmt.Println(perusahaan)
}