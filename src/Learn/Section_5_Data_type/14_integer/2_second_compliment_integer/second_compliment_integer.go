package main

import (
	"fmt"
	"strconv"
)

// why it dont show in 2 compliment
func main() {
	x := -5
	fmt.Printf("%b\n", x)
	fmt.Printf(strconv.FormatInt(-5, 2))
	fmt.Println(bitInt(int8(x)))
}

func bitInt(x int8) [8]byte {
	var result [8]byte // 1111 1011
	for i := 0 ; i < len(result) ; i++ {
		result[7 - i] = byte(x & 1)	
		x = x >> 1
	}
	return result
}
