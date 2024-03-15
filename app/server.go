package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading", err)
		return
	}

	request := string(buffer[:n])
	log.Printf("Received: %s", request)

	reqLines := strings.Split(request, "\r\n")
	path := strings.Split(reqLines[0], " ")[1]

	if path == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\n"))
	}

}
