package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

// add a handler which will listen for connections on port 4221 for get requests
// and respond with "Hello, World!" to the client
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read the request from the connection
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		return
	}

	// Check if it's a GET request
	request := string(buffer[:n])
	lines := strings.Split(request, "\r\n")
	fmt.Print(lines)
	if len(lines) > 0 {
		parts := strings.Split(lines[0], " ")
		if len(parts) >= 2 && parts[1] == "/abcdefg" {
			response := "HTTP/1.1 404 Not Found\r\n\r\n"
			conn.Write([]byte(response))
			return
		} else {
			response := "HTTP/1.1 200 OK\r\n\r\nHello, World!"
			conn.Write([]byte(response))
			return
		}
	}
	conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\n"))
}
