package main

import "fmt"

func main() {

	// Case 1 : Integer to String
	ex1 := string(0x265e)
	fmt.Println("case 1 : Integer -> string")
	for i := 0; i < len(ex1); i++ {
		fmt.Printf("%x ", ex1[i])
	}
	fmt.Println("\n\xe2\x99\x9e")
	fmt.Println("\xe2\x99\x9e" == ex1)

	// Case 2 : slice of byte to string
	fmt.Println("\nCase 2 : []byte -> string")
	ex2 := []byte{0xe2, 0x99, 0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String)

	// Case 3 : slice of rune to string
	fmt.Println("\nCase 3 : slice of rune -> string")
	ex3 := []rune{0x2654, 0x2655, 0x2656, 0x2657, 0x2658, 0x2659}
	fmt.Println(string(ex3))

	// Case 4 : string to slice of byte
	fmt.Println("\nCase 4 : string -> []byte")
	ex4 := []byte("Hello♕")
	fmt.Println(ex4)

	// Case 5 : string to slice of runes
	fmt.Println("\nCase 5 : string -> runes")
	ex5 := []rune("Hello♕")
	fmt.Println(ex5)
}
