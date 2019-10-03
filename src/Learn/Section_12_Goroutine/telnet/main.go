package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	log.SetFlags(log.Ltime)
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	// copy conn --> stdout
	log.Println("copy from conn to stdout")
	go io.Copy(os.Stdout, conn)
	log.Println("copy from stdin to conn")
	io.Copy(conn, os.Stdin)

}
