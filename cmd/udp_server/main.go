package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":9001")
	conn, _ := net.ListenUDP("udp", addr)
	buf := make([]byte, 1024)
	fmt.Println("UDP server listening on 9001")

	for {
		n, clientAddr, _ := conn.ReadFromUDP(buf)
		fmt.Println("client says:", string(buf[:n]))
		conn.WriteToUDP([]byte("hello client"), clientAddr)
	}
}
