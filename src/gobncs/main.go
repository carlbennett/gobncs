package gobncs

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Received: %s\n", string(buf[:n]))

		_, err = conn.Write(WriteMessage(0x25, []byte("\xBA\xAD\xF0\x0D")))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}
