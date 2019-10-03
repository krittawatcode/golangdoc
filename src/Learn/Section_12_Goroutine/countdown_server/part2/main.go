package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		log.Println("New client is connected : ", conn.RemoteAddr())
		if err != nil {
			log.Println(err)
			return
		}

		go countingDownHandler(conn)
	}
}

func countingDownHandler(conn net.Conn) {

	defer func() {
		io.WriteString(conn, "Your connection will be closed by server : ")
		conn.Close()
	}()

	io.WriteString(conn, "Enter number : ")
	input := bufio.NewScanner(conn)
	count := Scan(input)

	if count > 10 {
		return
	}

	for {
		io.WriteString(conn, strconv.Itoa(count)+"\n") // convert int from server to text and send back to client
		time.Sleep(time.Second)
		count--
		if count < 0 {
			io.WriteString(conn, "Enter number : ")
			count = Scan(input)
			if count == 0 {
				break
			}
		}
	}

	io.WriteString(conn, "Hello\n")
}

func Scan(input *bufio.Scanner) int {
	if ok := input.Scan(); !ok {
		log.Println("cannot scan value from conn")
		log.Println("Connection is closed by client")
		return 0
	}

	count, err := strconv.Atoi(input.Text()) // convert text from client to int
	if err != nil {
		log.Println("connot convert value from Text to int", input.Text())
		return 0
	}

	return count
}
