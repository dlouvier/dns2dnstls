package main

const (
	// HostListenAddress - Listen address to export the service
	HostListenAddress string = "0.0.0.0"
	// HostPort - Port where the service is exposing
	HostPort int = 8181
	// ForwarderDNS - External DNS resolver
	ForwarderDNS = "109.69.8.51:53" // 1.0.0.1 as my ISP (o2) blocks 1.1.1.1 and 1.0.0.0 addresses
)

func main() {
	go UDPServer()
	TCPServer()
}
