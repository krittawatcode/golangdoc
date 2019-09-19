package main

import (
	"fmt"
	"io"
	"net/http"
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

	var w io.Writer // Write(p []byte)(n int, err error)
	w = os.Stdout   // change dynamic type to *os.File

	// result := w.(MyIO) // panic
	result := w.(http.Handler)
	fmt.Printf("%T, %#v\n", w, w)
	fmt.Printf("%T, %#v\n", result, result)

}
