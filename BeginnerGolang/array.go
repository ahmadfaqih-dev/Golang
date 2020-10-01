package main

import "fmt"

func main () {
	var names [3]string

	// ubah data string
	names[0] = "Ahmad"
	names[1] = "Faqih"
	names[2] = "Arin"

	//len (cek panjang data array)
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, value := range names {
		//fmt.Println("index", index, "=", value)
		fmt.Println(value)

	}




	//fmt.Println(names) ambil semua data array
	//fmt.Println(names[0]) // ambil data satu2
	//fmt.Println(names[1])
	//fmt.Println(names[2])
}
