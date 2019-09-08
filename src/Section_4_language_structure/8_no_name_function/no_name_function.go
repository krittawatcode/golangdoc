package main

import "fmt"

func newIntPointer() *int {
	var x int
	return &x
}

func main() {
	// x := newIntPointer()
	fmt.Println(newIntPointer() == newIntPointer()) // compare in and out function -> return false bcz it have difference pointer
	fmt.Println(new(int) == new(int))               // new(int) mean new function with return pointer
}
