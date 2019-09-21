package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Ltime)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
		return
	}

	countingDownHandler(conn)
}

func countingDownHandler(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Hello")
}
