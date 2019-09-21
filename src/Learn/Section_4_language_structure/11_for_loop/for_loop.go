package main

import "fmt"

func main() {

	// for loop syntax
	/*
		for initialization ; condition ; post {
			body
		}
	*/

	// infinite loop
	/*
		for {
			fmt.Println("infinite loop")
			fmt.Println(time.Now().Clock()) // time syntax
		}
	*/

	// for i := 1; i < 5; i++ {
	// 	fmt.Println("still lower than 5")
	// }

	// for i := 1; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i == 3 {
	// 		fmt.Println("I about to terminate")
	// 		break // break from for loop
	// 	}
	// }

	for i := 1; i < 5; i++ {
		if i == 2 {
			continue // skip this condition
		}
		fmt.Println(i)
	}
}
