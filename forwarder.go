package main

import (
	"crypto/tls"
)

// QueryForwarder - TCP Client which connect to the external DNS service via TCP over TLS
func QueryForwarder(q []byte) []byte {
	// Create the connection
	c, err := tls.Dial("tcp", ForwarderDNS, &tls.Config{})
	Debugger(err)

	// Send the request to the forwarder
	_, err = c.Write(q)
	Debugger(err)

	buffer := make([]byte, 512)
	_, err = c.Read(buffer)
	//Debugger(err)

	// Ensure we close the client to the remote forwader
	defer c.Close()

	return buffer
}
