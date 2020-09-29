package main

import (
	"fmt"
	"runtime"
)

func main () {

	for i := 1; i <= 10; i++ {

		//if i == 1 {
		//	fmt.Println("Satu")
		//} else if i == 2 {
		//	fmt.Println("Dua")
		//} else if i == 3 {
		//	fmt.Println("Tiga")
		//} else {
		//	fmt.Println("Tidak tau")
		//}

		switch i {
		case 1:
			fmt.Println("Satu")
		case 2:
			fmt.Println("Dua")
		case 3:
			fmt.Println("Tiga")
		default:
			fmt.Println("Tidak tau")
		}
	
	}

	sistemOperasi := runtime.GOOS

	switch sistemOperasi {
	case "faqih":
		fmt.Println("Windows")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("tidak tau")

	}
}
