package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading", err)
			return
		}

		received := string(buffer[:n])
		log.Printf("Received: %s", received)
		if errors.Is(err, io.EOF) {
			return
		}
		conn.Write([]byte("HTTP/1.1 200 OK" + "\r\n\r\n"))
	}
}
