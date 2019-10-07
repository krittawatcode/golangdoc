package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	chSquare := make(chan int)
	done := make(chan string)

	go sender(ch, done)
	go square(ch, chSquare, done)
	go receiver(chSquare, done)

	for v := range done {
		fmt.Println("Done status : ", v)
	}

	fmt.Println("Main exit")

}

func sender(ch chan<- int, done chan<- string) {
	for i := 0; i <= 5; i++ {
		fmt.Println("Sending value : ", i)
		ch <- i
	}
	close(ch)
	fmt.Println("Sender is about to complete")
	done <- "Done from Sender"
	fmt.Println("Sender done")
}

func receiver(ch <-chan int, done chan<- string) {
	for v := range ch {
		fmt.Println("\tReceive value : ", v)
	}
	done <- "Done from Receiver"
	time.Sleep(100 * time.Millisecond)
	close(done)
}

func square(ch <-chan int, chSquare chan<- int, done chan<- string) {
	for v := range ch {
		chSquare <- v * v
	}

	close(chSquare)
	done <- "Done from Square"
}
