package main

import (
	"fmt"
	"io"
	"os"
)

type MyIO struct {
}

func (m MyIO) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m MyIO) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func main() {
	// x.(T)
	var w io.Writer // Write(p []byte)(n int, err error)
	w = os.Stdout

	result := w.(io.ReadWriter)
	fmt.Printf("%T, %#v\n", w, w)
	fmt.Printf("%T, %#v\n", result, result)

}
