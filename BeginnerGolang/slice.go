package main

import "fmt"

func main() {
	names := [5]string {         //penulisan array dengan cara lebih simple
		"Ahmad",
		"Faqih",
		"Arin",
		"Risma",
		"Rima",
}
	fmt.Println(names)

	slice1 := names[1:3]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)

	names[1] = "Muhammad"
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice1[1] = "Naza"
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

// slice: potongan data (bisa ambil potongan data dari array)
// slice: pointer penunjuk array