package main

import (
	"fmt"
	"unicode/utf8"
	// "reflect"
)

func main() {
	x := "ทดสอบ"
	for i, v := range x {
		// fmt.Println(i, v, reflect.TypeOf(v))
		fmt.Printf("%d, %c\n", i, v)
	}

	fmt.Println("utf8.RuneCountInString : ", utf8.RuneCountInString(x))
	// fmt.Println("utf8.DecodeRuneInString : ", utf8.DecodeRuneInString(x[3:7]))

	/* ----- Part II -----*/

	for i := 0; i < len(x); {
		r, s := utf8.DecodeRuneInString(x[i:])
		i += s
		fmt.Printf("%c-", r)
	}

}
