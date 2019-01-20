package main

import (
	"net"
)

// TCPRequestHandler - Deal with the TCP Request
func TCPRequestHandler(c net.Conn) {
	buffer := make([]byte, 512)
	_, err := c.Read(buffer)
	Debugger(err)

	ans := QueryForwarder(buffer)
	_, err = c.Write(ans)
	Debugger(err)

	defer c.Close()
}
