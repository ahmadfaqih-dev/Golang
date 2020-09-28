package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int32 = 10
	var y int64 = int64(x)

	fmt.Println(y)

	var z float64 = float64(y)

	fmt.Println(z)

	var b float64 = 3.3

	var a int32 = int32(b)

	fmt.Println(a)

//	conversi string ke integer

	var nilai string = "100"

	nilaiInt, _ := strconv.Atoi(nilai)

	fmt.Println(nilaiInt)

//	conversion integer to string

	nilaiString := strconv.Itoa(-100)

	fmt.Println(nilaiString)

	nilaiBool, _ := strconv.ParseBool("true")

	fmt.Println(nilaiBool)
}
