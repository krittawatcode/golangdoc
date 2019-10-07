package main

import "fmt"

func main() {

	ch := make(chan int)
	done := make(chan string) // done <- struct{}{}
	go sender(ch, done)
	go receiver(ch, done)
	doneStatus := <-done
	fmt.Println("Done status : ", doneStatus)
	doneStatus = <-done
	fmt.Println("Done status : ", doneStatus)
	fmt.Println("Main exit")

}

func sender(ch chan int, done chan string) {
	for i := 0; i <= 5; i++ {
		fmt.Println("Sending value : ", i)
		ch <- i
	}
	close(ch)
	fmt.Println("Sender is about to complete")
	done <- "Done from Sender"
	fmt.Println("Sender done")
}

func receiver(ch chan int, done chan string) {
	fmt.Println("\tReceive Value is : ", <-ch) // 0
	fmt.Println("\tReceive Value is : ", <-ch) // 1
	fmt.Println("\tReceive Value is : ", <-ch) // 2
	fmt.Println("\tReceive Value is : ", <-ch) // 3
	fmt.Println("\tReceive Value is : ", <-ch) // 4
	fmt.Println("\tReceive Value is : ", <-ch) // 5
	fmt.Println("\tReceive Value is : ", <-ch) // 6 <- deadlock! if not close
	// if use close() -> receiver will receive only zero value of type of channels
	done <- "Done from Receiver"
}
