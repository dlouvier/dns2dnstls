package main

import (
	"crypto/tls"
	"fmt"
)

// QueryForwarder - TCP Client which connect to the external DNS service via TCP over TLS
func QueryForwarder(q []byte) []byte {
	fmt.Println("Aqui estoy yo")
	// Create the connection

	tr := &tls.Config{
		InsecureSkipVerify: true}

	c, err := tls.Dial("tcp", ForwarderDNS, tr)

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
