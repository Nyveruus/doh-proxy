// Listen UDP 53, parse DNS query, dispatch DoH request https tcp 443, receive reply, parse response to plain dns maintain ECH, response to client UDP
// goroutine every query

package main

import (
	"fmt"
	"os"
	"net"
	"net/http"
	"io"
	"bytes"
	"codeberg.org/miekg/dns"
)

const listenAddr string = ":53"
const bufferSize int = 512
const DoHAddr string = "https://cloudflare-dns.com/dns-query"

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
	if err := new(dns.Msg).Unpack(query); err != nil {
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
	if _, err := connection.WriteTo(response, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write response: %v", err)
	}
}

func forwardDoH(query []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", DoHAddr, bytes.NewReader(query))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/dns-message")
	request.Header.Set("Accept", "application/dns-message")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("upstream returned %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}
