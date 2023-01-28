package server

import (
	"log"
	"net"

	"github.com/carlbennett/gobncs/message"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	defer log.Printf("(%s) Connection terminated\n", conn.RemoteAddr())

	log.Printf("(%s) Connection established\n", conn.RemoteAddr())

	protocol, err := message.ReadProtocolByte(conn)
	if err != nil {
		log.Printf("(%s) Error with receiving protocol type from socket", conn.RemoteAddr())
		return
	}

	if protocol != 0x01 {
		log.Printf("(%s) Unknown protocol type (0x%02x) requested", conn.RemoteAddr(), protocol)
		return
	}

	for {
		message, err := message.ReadMessage(conn)
		if message == nil || err != nil {
			return
		}
		go HandleMessage(message)
	}
}

func HandleMessage(message *message.Message) {
	// Do something with the message here
}
