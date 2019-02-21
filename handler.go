package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	name       string
	reader     *bufio.Scanner
	writer     *bufio.Writer
	connection net.Conn
}

// TCPRequestHandler - Deal with the TCP Request
func TCPRequestHandler(c net.Conn) {
	defer c.Close()
	log.Println("TCP Request received")

	client := &Client{
		reader:     bufio.NewScanner(c),
		writer:     bufio.NewWriter(c),
		connection: c,
	}

	for {
		if ok := client.reader.Scan(); !ok {
			break
		}
		fmt.Println(client.reader.Bytes())
	}

	// // client := &Client{
	// // 	reader:     bufio.NewReader(c),
	// // 	writer:     bufio.NewWriter(c),
	// // 	connection: c,
	// // }
	// blabla, _ := client.reader.Re
	// fmt.Println(blabla)

	// //client.name = strings.TrimSpace(client.name)

	// //lobby = append(lobby, *client)
	// fmt.Println("Client connected: " + client.name)
}
