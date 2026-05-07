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

func main() {
	connection, err := net.ListenPacket(udp, listenAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v", err)
		os.Exit(1)
	}
}
