package main

import "fmt"

func main() {
	x := int8(127)
	y := int8(1)
	fmt.Println(x + y) // -128

	p := 127
	q := 2
	fmt.Println(int8(p + q)) // -127

	// convert float to int8
	z := 3.999
	fmt.Println(int8(z)) // 3
}
