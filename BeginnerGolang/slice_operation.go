package main

import "fmt"

func main () {
//	membuat slice dengan slice
//	slice1 := []string{
//		"Ahmad",
//		"Faqih",
//		"nazafarin",
//	}
	slice1 := make([]string, 3)
	slice1[0] = "Ahmad"
	slice1[1] = "Faqih"
	slice1[2] = "nazafarin"

	//	append adalah menambahkan data slice
	 slice2 := append(slice1, "toni", "roni")

	 fmt.Println(slice1)
	 fmt.Println(slice2)

	 // make adalah tipe data slice
	 slice3 := make([]string, 3)

	 //copy adalah untuk mengcopy data slice
	 copy(slice3, slice1)

	 fmt.Println(slice1)
	 fmt.Println(slice3)

	 slice1[0] = "Muhammad"

	fmt.Println(slice1)
	fmt.Println(slice3)

}
