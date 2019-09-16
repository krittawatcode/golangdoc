package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// x.(T)
	var w io.Writer // Write(p []byte)(n int, err error)
	w = os.Stdout
	result := w.(*os.File)
	fmt.Printf("%T, %#v", result, result)
}
