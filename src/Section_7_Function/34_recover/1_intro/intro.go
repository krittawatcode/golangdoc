package main

import (
	"fmt"
	"os"
)

var cygwinPath = os.Getenv("CYGWIN")

func A(x int) {
	defer func() {
		fmt.Println("Defer in A")
	}()
	B(x)
	fmt.Println("Hello in A")
}

func B(x int) {
	defer func() {
		fmt.Println("Defer in B")
	}()
	C(x)
	fmt.Println("Hello in B")
}

func C(x int) {
	defer func() {
		fmt.Println("Defer in C")
		if p := recover(); p != nil {
			fmt.Println("handling panic ")
		}
	}()
	if x == 4 {
		panic("for no reason")
	}
}

func main() {
	A(4)
}
