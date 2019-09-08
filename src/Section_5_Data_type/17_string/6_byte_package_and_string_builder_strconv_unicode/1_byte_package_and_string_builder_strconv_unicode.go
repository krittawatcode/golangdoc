package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	buff := bytes.Buffer{}

	fmt.Println("----- part 1 : buff string -----")

	buff.WriteRune('y')
	buff.WriteRune('o')
	fmt.Println(buff.String())

	fmt.Println("----- part 2 : string builder -----")
	buff2 := strings.Builder{}
	buff2.WriteRune('y')
	buff2.WriteRune('o')
	fmt.Println(buff2.String())

	fmt.Println("----- part 3 : strconv -----")
	atoi, _ := strconv.Atoi("123")
	itoa := strconv.Itoa(123)
	fmt.Println(atoi, reflect.TypeOf(atoi))
	fmt.Println(itoa, reflect.TypeOf(itoa))

	// strconv.ParseBool
	fmt.Println("----- part 3_2 : strconv bool -----")
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("False"))
	fmt.Println(strconv.ParseBool("fALse"))
	fmt.Println(strconv.ParseBool("0"))

	// unicode package
	fmt.Println("----- part 4 : unicode package -----")
	// look in video

}
