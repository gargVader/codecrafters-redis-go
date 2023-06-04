package main

import (
	"fmt"
	"net"
	"os"
)

const (
	HOST = "127.0.0.1"
	PORT = "6379"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		fmt.Println("Failed to bind to port "+PORT+" ", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for {
		connection, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		handleRequest(connection)
	}
}

func handleRequest(connection net.Conn) {
	buffer := make([]byte, 1024)

	_, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Println(string(buffer))
	connection.Write([]byte("+PONG\r\n"))
	// Close
	connection.Close()

}
