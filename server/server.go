package server

import (
	"fmt"
	"log"
	"net"

	"github.com/carlbennett/gobncs/message"
	"github.com/carlbennett/gobncs/parser"
)

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	log.Printf("(%s) connection established; waiting for protocol type request\n", remoteAddr)
	defer log.Printf("(%s) connection terminated\n", remoteAddr)

	protocol, err := message.ReadProtocolByte(conn)
	if err != nil {
		return err
	}

	switch protocol {
	case 0x01, 0x02, 0x03, 0x63:
		log.Printf("(%s) protocol type (0x%02X) requested", remoteAddr, protocol)
	default:
		log.Printf("(%s) unknown protocol type (0x%02X) requested; terminating connection", remoteAddr, protocol)
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

func HandleMessage(conn net.Conn, messageData *message.Message) {
	remoteAddr := conn.RemoteAddr()
	messageId := messageData.ID
	messageName := message.MessageIdToName(messageId)

	log.Printf("(%s) message ([0x%02X] %s) received from client; parsing", remoteAddr, messageId, messageName)

	var err error
	switch messageId {
	case message.SID_NULL:
		err = parser.ParseSID_NULL(conn, messageData)
	case message.SID_PING:
		err = parser.ParseSID_PING(conn, messageData)
	case message.SID_AUTH_INFO:
		err = parser.ParseSID_AUTH_INFO(conn, messageData)
	default:
		err = fmt.Errorf("unknown message id (0x%02X); terminating connection", messageId)
	}

	if err != nil {
		log.Printf("(%s) error parsing message ([0x%02X] %s): %s", remoteAddr, messageId, messageName, err)
		conn.Close()
	}
}
