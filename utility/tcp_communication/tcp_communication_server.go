package tcp_communication

import (
	"fmt"
	"net"
)

func TCPServer() {
	listener, _ := net.Listen("tcp", ":9000")
	fmt.Println("server listening on 9000")

	for {
		conn, _ := listener.Accept()
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("client says:", string(buf[:n]))
	conn.Write([]byte("hello client"))
}

// Client
//   │
//   │ 1. Connect
//   ▼
// Server
//   │
//   │ 2. Send / Receive data
//   ▼
// Client
//   │
//   │ 3. Close connection
//   ▼
// Server

// 1. Connect (establish connection)
// Client phải mở connection tới server bằng:
// - IP + Port = 127.0.0.1:8080
// - Ở tầng TCP sẽ diễn ra: SYN / SYN-ACK / ACK (Synchronize / Acknowledgement)
// => connection established

// 2. Send / Receive data (Sau khi connect)
// - Client có thể: Read response + Write data (bytes)

// 3. Close connection (when done): FIN (Finish) / ACK
