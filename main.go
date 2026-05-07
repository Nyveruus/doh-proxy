// Listen UDP 53, parse DNS query, dispatch DoH request https tcp 443, receive reply, parse response to plain dns maintain ECH, response to client UDP
// goroutine every query

package main

import (
	"fmt"
	"os"
	"net"
	"codeberg.org/miekg/dns"
)

const listenAddr string = ":53"
const bufferSize uint16 = 512

func main() {
	connection, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v", err)
		os.Exit(1)
	}

	buffer [512]byte
	for {
		bytesRead, addr, err := connection.ReadFrom(buffer[:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read error: %v", err)
			continue
		}
	}
}
