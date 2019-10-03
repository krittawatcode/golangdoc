package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// conn, err := listener.Accept()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// countingDownHandler(conn)

	for {
		conn, err := listener.Accept()
		log.Println("New client is connected : ")
		if err != nil {
			log.Println(err)
			return
		}

		countingDownHandler(conn)
	}
}

func countingDownHandler(conn net.Conn) {
	// defer conn.Close()
	defer func() {
		io.WriteString(conn, "Your connection will be closed by server : ")
		conn.Close()
	}()

	count := 5

	for {
		io.WriteString(conn, "Hello savior\n")
		time.Sleep(time.Second)
		count--
		if count == 0 {
			io.WriteString(conn, "Enter number : ")
			// todo
			break
		}
	}

	io.WriteString(conn, "Hello\n")
}
