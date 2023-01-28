package server

import (
	"log"
	"net"

	"github.com/carlbennett/gobncs/message"
)

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	log.Printf("(%s) Connection established; waiting for protocol type request\n", remoteAddr)
	defer log.Printf("(%s) Connection terminated\n", remoteAddr)

	protocol, err := message.ReadProtocolByte(conn)
	if err != nil {
		return err
	}

	switch protocol {
	case 0x01, 0x02, 0x03, 0x63:
		log.Printf("(%s) Protocol type (0x%02X) requested", remoteAddr, protocol)
	default:
		log.Printf("(%s) Unknown protocol type (0x%02X) requested; terminating connection", remoteAddr, protocol)
		return err
	}

	for {
		message, err := message.ReadMessage(conn)
		if message == nil || err != nil {
			return err
		}
		go HandleMessage(conn, message)
	}
}

// refactor this file for me please 

func HandleMessage(conn net.Conn, messageData *message.Message) {
	remoteAddr := conn.RemoteAddr()
	messageId := messageData.ID

	log.Printf("(%s) Message id (0x%02X: %s) received from client; parsing", remoteAddr, messageId, message.MessageIdToName(messageId))

	switch messageId {
	case message.SID_NULL:
		message.ParseMessage(conn, messageData)
	case message.SID_PING:
		message.ParseMessage(conn, messageData)
	case message.SID_AUTH_INFO:
		message.ParseMessage(conn, messageData)
	default:
		log.Printf("(%s) Unknown message id (0x%02X); terminating connection", remoteAddr, messageId)
		conn.Close()
	}
}
