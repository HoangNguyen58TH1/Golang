package tcp_communication

import (
	"fmt"
	"net"
)

func TCPClient() {
	conn, _ := net.Dial("tcp", "localhost:9000")
	conn.Write([]byte("hello server"))
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("server reply:", string(buf[:n]))
	conn.Close()
}

// Client                    Server
//   │                         │
//   │  Dial()                 │
//   │──────────── connect ───►│
//   │                         │
//   │ Write("hello server")   │
//   │────────────────────────►│
//   │                         │
//   │                 Read()  │
//   │                         │
//   │                 Write() │
//   │◄────────────────────────│
//   │ Read()                  │
//   │                         │
//   │ Close()                 │
