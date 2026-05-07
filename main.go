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
const bufferSize int = 512

func main() {
	connection, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v", err)
		os.Exit(1)
	}

	buffer := make([]byte, bufferSize)
	for {
		bytesRead, addr, err := connection.ReadFrom(buffer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read error: %v", err)
			continue
		}
		query := make([]byte, bytesRead)
		copy(query, buffer)

		go handleQuery(connection, addr, query)
	}
}

func handleQuery(connection net.PacketConn, addr net.Addr, query []byte) {
	//check if valid
	m := new(dns.Msg)
	if err := m.Unpack(query); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unpack: %v", err)
		return
	}

	//doh
	response, err := forwardDoH(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "DoH query failed: %v", err)
		return
	}
	//write back to client
}

func forwardDoH(query []byte) {

}
