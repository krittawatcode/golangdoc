package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	s := a[1:3]
	ss := make([]int, 5, 10) // make([]int, len, cap) or make([]int, len)
	ss[6] = 99
	fmt.Println(len(s), cap(s), s)
	fmt.Println(len(ss), cap(ss), ss)
	// panic: runtime error: index out of range
}
