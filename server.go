package main

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"

	"golang.org/x/net/dns/dnsmessage"
)

// TCPServer - Start a server listening under TCP Protocol
func TCPServer() {
	srv, err := net.Listen("tcp", HostListenAddress+":"+strconv.Itoa(HostPort))
	Debugger(err)
	// Close the listener when the application closes.
	defer srv.Close()

	log.Println("DNS Proxy server started on " + HostListenAddress + ":" + strconv.Itoa(HostPort) + " (TCP)")

	for {
		// Allow concurrency of multiple connections
		conn, err := srv.Accept()

		Debugger(err)
		// Handle the request and sets it is an TCP request
		go TCPRequestHandler(conn)
	}
}

// UDPServer - Start a server listening on UDP Protocol
func UDPServer() {
	srv, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP((HostListenAddress)), Port: HostPort})
	Debugger(err)

	// Close the listener when the application closes.
	defer srv.Close()

	log.Println("DNS Proxy server started on " + HostListenAddress + ":" + strconv.Itoa(HostPort) + " (UDP)")

	for {
		buffer := make([]byte, 512)
		PkgSizeInt, addr, err := srv.ReadFromUDP(buffer)
		Debugger(err)
		log.Println("UDP Request received")

		// I need to create an TCP package, which requires to know the size of the package
		PkgSize := make([]byte, 2)
		binary.BigEndian.PutUint16(PkgSize, uint16(PkgSizeInt))

		buffer = append(PkgSize, buffer[:510]...)

		//Query the external DNS resolver
		ans := QueryForwarder(buffer)

		var m dnsmessage.Message
		err = m.Unpack(ans[2:])
		Debugger(err)

		ans, _ = m.Pack()

		_, err = srv.WriteToUDP(ans, addr)
	}
}
