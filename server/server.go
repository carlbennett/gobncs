package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"

	"github.com/carlbennett/gobncs/clientstate"
	"github.com/carlbennett/gobncs/message"
	"github.com/carlbennett/gobncs/parser"
)

func HandleConnection(conn net.Conn) error {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr()
	log.Printf("(%s) connection established; waiting for protocol type request\n", remoteAddr)
	defer log.Printf("(%s) connection terminated\n", remoteAddr)

	state := &clientstate.ClientState{
		Conn:         conn,
		Ping:         -1,
		PingCookie:   rand.Uint32(),
		Platform:     clientstate.PLATFORM_ZERO,
		Product:      clientstate.PRODUCT_ZERO,
		RemoteAddr:   remoteAddr,
		TimezoneBias: 0,
	}
	clientstate.AddClientState(conn, state)
	defer clientstate.RemoveClientState(conn)

	protocol, err := clientstate.ReadProtocolType(conn)
	if err != nil {
		return err
	}
	state.ProtocolType = protocol

	switch protocol {
	case 0x01:
		log.Printf("(%s) protocol type (0x%02X) requested", remoteAddr, protocol)
	default:
		log.Printf("(%s) unknown protocol type (0x%02X) requested; terminating connection", remoteAddr, protocol)
		return err
	}

	// begin game protocol message stream
	for {
		message, err := message.ReadMessage(conn)
		if message == nil || err != nil {
			return err
		}
		go HandleMessage(state, message)
	}
}

func HandleMessage(state *clientstate.ClientState, messageData *message.Message) {
	remoteAddr := state.RemoteAddr
	messageId := messageData.ID
	messageName := message.MessageIdToName(messageId)

	log.Printf("(%s) message ([0x%02X] %s) received from client; parsing", remoteAddr, messageId, messageName)

	var err error
	switch messageId {
	case message.SID_NULL:
		err = parser.ParseSID_NULL(state, messageData)
	case message.SID_PING:
		err = parser.ParseSID_PING(state, messageData)
	case message.SID_AUTH_INFO:
		err = parser.ParseSID_AUTH_INFO(state, messageData)
	default:
		err = fmt.Errorf("unknown message id (0x%02X); terminating connection", messageId)
	}

	if err != nil {
		log.Printf("(%s) error parsing message ([0x%02X] %s): %s", remoteAddr, messageId, messageName, err)
		state.Conn.Close()
	}
}
