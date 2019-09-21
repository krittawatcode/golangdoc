package main

import "fmt"

func main() {

	// char in go don't must have only 1 byte

	x := "hi-สวัสดี" // []byte{0,1,2,...,21} -> slide of byte
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	fmt.Println(x, len(x)) // the length of string is string in byte

	/* ----- check x is base 16 or not ----- */
	// for i := 0; i < len(x); i++ {
	// 	fmt.Printf("%d = %x\n", i, x[i])
	// }

	/* ----- Create y in byte ----- */
	// for i := 0; i < len(x); i++ {
	// 	fmt.Printf("0x%x, ", x[i])
	// }

	fmt.Println(string(y[0])) // h
	fmt.Println(string(y[1])) // i
	fmt.Println(string(y[2])) // -
	fmt.Println(string(y[3])) // ส
	fmt.Println(string(y[4])) // ว

}
