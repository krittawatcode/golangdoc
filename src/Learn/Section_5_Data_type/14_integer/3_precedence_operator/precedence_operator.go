package main

import "fmt"

func main() {

	// % will retrun เครื่องหมาย + - by first number
	fmt.Println(-5 % -2) // -1

	// 1001 9
	// 1110 8 + 4 + 2 = 14
	// 1111 8+4+2+1 = 15
	fmt.Println(9 | 14)

	// 101 5
	// 010 2
	// 000 0
	fmt.Println(5 & 2) // and

	// 101 5
	// 101 ^2 = 5
	// 101 5
	fmt.Println(5 &^ 2) // and not
}
