package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	msg string
}

func (m MyError) Error() string {
	return m.msg
}

func main() {
	// ioutil.ReadFile()
	// os.Open()

	var err error
	err = errors.New("sdsdg")
	fmt.Printf("%T, %#v\n", err, err)

	fmt.Println(errors.New("sad") == errors.New("sad"))
	// will return false because errors.New will allowcate new space to hold value
	// it will have same value but not same address

	err = MyError{"this is an error"}
	fmt.Printf("%T ,%#v\n", err, err)
}
