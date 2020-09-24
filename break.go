package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		//if i % 2 == 0 {
		//	continue
		//}

		fmt.Println(i)

		if i == 20 {
			break
		}

	}
}


// continue adl skip program kode yang dibawahnya dan dia akan melanjutkan perulangan selanjutnya
// break menghentikan perulangan selanjutnya