package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Close the listener when listener close
	defer l.Close()

	fmt.Println("Listening on localhost:8080")
	for {
		//Listen for incoming connnections
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Write to connection
		io.WriteString(conn, "I see you connected")
		conn.Close()
	}
}
