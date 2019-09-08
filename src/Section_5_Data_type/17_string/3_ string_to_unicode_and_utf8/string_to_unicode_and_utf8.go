package main

import "fmt"

func main() {
	x := "Hi-สวัสดี"
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	fmt.Println(x, len(x))
	fmt.Println(string(y[0]))  // h
	fmt.Println(string(y[1]))  // i
	fmt.Println(string(y[2]))  // -
	fmt.Println(string(y[3]))  // ส
	fmt.Println(string(y[4]))  // ว
	fmt.Println(string(0xE2A)) // ส

	/* ----- Part II ----- */
	fmt.Println("--------- Part II ----------")

	z := []rune(x)
	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Println(len(z))
	fmt.Printf("%q", z[5])

}
