package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

// TCPRequestHandler - Deal with the TCP Request
func TCPRequestHandler(c net.Conn) {
	log.Println("TCP Request received")
	// receive the message

	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		// read client request data
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}

		fmt.Println(bytes)

		PkgSize := make([]byte, 2)
		binary.BigEndian.PutUint16(PkgSize, uint16(len(bytes)))

		bytes = append(PkgSize, bytes...)

		ans := QueryForwarder(bytes)
		_, err = c.Write(ans)
	}

	//defer c.Close()
}
