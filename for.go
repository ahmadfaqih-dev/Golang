package main

import "fmt"

func main() {

	// for i := 1; i <= 10; i++ (Lebih singkat)
	i := 1
	for i <= 10 {
		fmt.Println(i, "Hello word")
		i++
	}

	for i := 10; i >= 1; i-- {
		fmt.Println(i, "Ahmad Faqih")
	}

}
