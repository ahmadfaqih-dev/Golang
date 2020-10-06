package main

import (
	"fmt"
)

func main() {
	 mahasiswa := make(map[string]string)

	 mahasiswa["1001"] = "Ahmad"
	 mahasiswa["1002"] = "Faqih"
	 mahasiswa["1003"] = "Nazafarin"

	 fmt.Println(mahasiswa)
	 fmt.Println(mahasiswa["1001"])
	 fmt.Println(mahasiswa["1002"])
	 fmt.Println(mahasiswa["1003"])

	for nim, name := range mahasiswa {
		fmt.Println("Nim", nim, "Bernama", name)
	}
}
