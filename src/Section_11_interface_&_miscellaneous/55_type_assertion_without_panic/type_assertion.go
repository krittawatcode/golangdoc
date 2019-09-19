package main

import (
	"fmt"
	"io"
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

	var w io.Writer
	// w = os.Stdout
	w = MyIO{}

	result, ok := w.(io.Reader)
	fmt.Printf("%T, %#v\n", w, w)
	fmt.Printf("%T, %#v, %v\n", result, result, ok)

	Value("sad") // assert ok. value is :  sad
	Value(43)    // assert not ok

}

func Value(key interface{}) {
	if keyAsString, ok := key.(string); ok {
		fmt.Println("assert ok. value is : ", keyAsString)
		return
	}
	fmt.Println("assert not ok")
}
