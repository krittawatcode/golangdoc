package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var w io.Writer // Write(p []byte)(n int, err error)
	w = os.Stdout   // change dynamic type to os.Stdout

	result := w.(io.ReadWriter) // assert that concrete value stored in w is underlying type of io.ReadWriter -> still get *os.File bcz io.Reader doesn't have underlying type
	// result will have w value that have type *os.File and must implement method from both interface
	// https://stackoverflow.com/questions/38816843/explain-type-assertions-in-go
	/*
		assert that the dynamic type of w(os.Stdout) implements the interface io.ReadWriter.
		Return: The io.ReadWriter interface whose dynamic typ is *os.File (the one that w is holding)
	*/
	fmt.Printf("%T, %#v\n", w, w)
	fmt.Printf("%T, %#v\n", result, result)
	// result will get two method -> result.Read() result.Write()

}
