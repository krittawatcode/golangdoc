package main

import "fmt"

// send only
func sender(ch chan<- int) {
	ch <- 88
	// f(ch) <= cannot use it
}

func f(ch chan int) {
	ch <- 99
}

// receive only // collect int from 'ch chan'
func receiver(ch <-chan int, done chan bool) {
	fmt.Println("Receive : ", <-ch)
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sender(ch)
	go receiver(ch, done)
	<-done
}
