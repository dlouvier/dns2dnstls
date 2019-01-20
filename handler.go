package main

import (
	"log"
	"net"
)

// TCPRequestHandler - Deal with the TCP Request
func TCPRequestHandler(c net.Conn) {
	log.Println("TCP Request received")
	buffer := make([]byte, 512)
	_, err := c.Read(buffer)
	Debugger(err)

	ans := QueryForwarder(buffer)
	_, err = c.Write(ans)
	Debugger(err)

	defer c.Close()
}
