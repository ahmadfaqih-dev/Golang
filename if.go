package main

import "fmt"

func main () {

	for i := 1; i <= 30; i++ {
		//if i % 2 == 1{
		//	fmt.Println("Ganjil")
		//} else {
		//	fmt.Println(i)
		//}


		//	if i % 3 = Foo
		//	if i % 5 = Bar
		// 	if i % 3 && i % 5 = FooBar

		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FooBar")
		} else if i%3 == 0 {
			fmt.Println("Foo")
		} else if i % 5 == 0 {
			fmt.Println("Bar")
		} else {
			fmt.Println(i)
		}
	}


}
