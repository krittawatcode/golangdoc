package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// x.(T)
	/*
		x is an interface
		x.(T) is mean dynamic type of x is satisfy T or not ?
		dynamic type mean type of value in interface
		https://stackoverflow.com/questions/20518457/need-clarification-about-dynamic-types-in-golang?fbclid=IwAR2PWgf8BIBSZtqUrqg9MRUyd_xTFdqMSAtdymmWfUJ1J6XpAm8TS71TFHM
	*/
	var w io.Writer // Write(p []byte)(n int, err error)
	// w -> interface type -> Writer
	w = os.Stdout // interface dynamic value is type *File

	result := w.(*os.File) // check dynamic value is *os.File or not
	fmt.Printf("%T, %#v\n", w, w)
	fmt.Printf("%T, %#v\n", result, result)

	result2 := w.(*bytes.Buffer)
	fmt.Printf("%T, %#v\n", w, w)
	fmt.Printf("%T, %#v\n", result2, result2)
}
