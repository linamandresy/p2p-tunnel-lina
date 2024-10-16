package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var Clients = make(map[string]net.Conn)
var mu sync.Mutex

func main() {
	listenner, err := net.Listen("tcp", ":1309")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server started on port:1309")
	for {
		conn, err := listenner.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	ip := assignIP(conn)
	fmt.Printf("New Client connected with IP : %s\n", ip)

	mu.Lock()
	Clients[ip] = conn
	mu.Unlock()

	go forwarPacket(conn, ip)
}

func forwarPacket(conn net.Conn, ip string) {
	defer conn.Close()
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
		packet := buf[:n]
		destIp := extractIP(packet)
		mu.Lock()
		destConn, ok := Clients[destIp]
		if ok {
			destConn.Write(packet)
		}
		mu.Unlock()
	}
}

func extractIP(packet []byte) string {
	return strings.Split(string(packet), ":")[0]
}

func assignIP(conn net.Conn) string {
	ip := fmt.Sprintf("13.0.0.%d", len(Clients)+2)
	conn.Write([]byte(fmt.Sprintf("%s", ip)))
	return ip
}
