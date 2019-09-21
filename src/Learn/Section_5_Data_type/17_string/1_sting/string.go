package main

import "fmt"

func main() {
	eng := "Hello"
	th := "สวัสดี"
	fmt.Println(eng, len(eng), string(eng[0])) // length = 5  // eng[0] = H
	fmt.Println(th, len(th), string(th[0]))    // length = 18 ??? // th[0] = a ???

}
