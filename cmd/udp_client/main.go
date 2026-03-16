package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "localhost:9001")
	conn, _ := net.DialUDP("udp", nil, addr)
	conn.Write([]byte("hello server"))
	buf := make([]byte, 1024)
	n, _, _ := conn.ReadFromUDP(buf)
	fmt.Println("server reply:", string(buf[:n]))
	conn.Close()
}
