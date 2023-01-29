package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/carlbennett/gobncs/message"
)

func ParseSID_NULL(conn net.Conn, message *message.Message) error {
	if message.Length != 4 {
		return fmt.Errorf("invalid message length (expected 4, got %d)", message.Length)
	}

	return nil
}

func ParseSID_PING(conn net.Conn, message *message.Message) error {
	if message.Length != 8 {
		return fmt.Errorf("invalid message length (expected 8, got %d)", message.Length)
	}

	var cookie uint32
	err := binary.Read(bytes.NewReader(message.Body), binary.LittleEndian, &cookie)
	if err != nil {
		return fmt.Errorf("failed to read cookie from message body: %v", err)
	}

	log.Printf("cookie: %08X", cookie)
	return nil
}

func ParseSID_AUTH_INFO(conn net.Conn, message *message.Message) error {
	if message.Length < 38 {
		return fmt.Errorf("invalid message length (expected at least 38, got %d)", message.Length)
	}

	/** Client->Server Format:
	 * (UINT32) Protocol ID
	 * (UINT32) Platform code
	 * (UINT32) Product code
	 * (UINT32) Version byte
	 * (UINT32) Language code
	 * (UINT32) Local IP
	 * (UINT32) Time zone bias
	 * (UINT32) MPQ locale ID
	 * (UINT32) User language ID
	 * (STRING) Country abbreviation
	 * (STRING) Country
	 */

	return nil
}

func WriteSID(message *message.Message, id message.MessageId, length uint16) error {
	message.ID = id
	message.Length = 4 + length
	return nil
}

func WriteSID_PING(reply *message.Message, cookie uint32) error {
	err := WriteSID(reply, message.SID_PING, 4)
	if err != nil {
		return err
	}

	buffer := &bytes.Buffer{}
	err = binary.Write(buffer, binary.LittleEndian, &cookie)
	if err != nil {
		return err
	}

	reply.Body = buffer.Bytes()
	return nil
}
