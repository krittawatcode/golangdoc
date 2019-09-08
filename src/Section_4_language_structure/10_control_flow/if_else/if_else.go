package main

import "fmt"

func main() {

	// if statement (condition must be true to run func inside)
	// if condition {
	// 	fmt.Println("go!")
	// }

	if false {
		fmt.Println("go!")
	} else {
		fmt.Println("Go")
	}

	///////////////////////////////////////

	x := 5
	if v := x + 1; (x <= 5) || false { // v can use only inside this flow
		fmt.Println("Inside if", v)
	} else {
		fmt.Println("Inside else", v)
	}

}
