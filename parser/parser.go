package parser

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"

	"github.com/carlbennett/gobncs/clientstate"
	"github.com/carlbennett/gobncs/message"
)

func ParseSID_NULL(state *clientstate.ClientState, message *message.Message) error {
	if message.Length != 4 {
		return fmt.Errorf("invalid message length (expected 4, got %d)", message.Length)
	}

	return nil
}

func ParseSID_PING(state *clientstate.ClientState, message *message.Message) error {
	if message.Length != 8 {
		return fmt.Errorf("invalid message length (expected 8, got %d)", message.Length)
	}

	/** Client<->Server Format:
	 * (UINT32) Ping Cookie
	 */

	var cookie uint32
	err := binary.Read(bytes.NewReader(message.Body), binary.LittleEndian, &cookie)
	if err != nil {
		return fmt.Errorf("failed to read cookie from message body: %v", err)
	}

	if state.PingCookie != cookie {
		log.Printf("(%s) stale ping cookie; rejecting late SID_PING response", state.RemoteAddr)
		return nil
	}

	state.PingCookie = rand.Uint32() // change cookie so that repeated reply is considered stale
	return nil
}

func ParseSID_AUTH_INFO(state *clientstate.ClientState, payload *message.Message) error {
	if payload.Length < 38 {
		return fmt.Errorf("invalid message length (expected at least 38, got %d)", payload.Length)
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

	reader := bytes.NewReader(payload.Body)

	var protocolId uint32
	err := binary.Read(reader, binary.LittleEndian, &protocolId)
	if err != nil {
		return fmt.Errorf("failed to read protocol id: %v", err)
	}
	if protocolId != 0 {
		return fmt.Errorf("unknown protocol id (expected 0, got %d)", protocolId)
	}

	var platformId uint32
	err = binary.Read(reader, binary.LittleEndian, &platformId)
	if err != nil {
		return fmt.Errorf("failed to read platform id: %v", err)
	}
	state.Platform = clientstate.Platform(platformId)

	var productId uint32
	err = binary.Read(reader, binary.LittleEndian, &productId)
	if err != nil {
		return fmt.Errorf("failed to read platform id: %v", err)
	}
	state.Product = clientstate.Product(productId)

	var versionId uint32
	err = binary.Read(reader, binary.LittleEndian, &versionId)
	if err != nil {
		return fmt.Errorf("failed to read version id: %v", err)
	}
	state.VersionId = versionId

	var languageCode uint32
	err = binary.Read(reader, binary.LittleEndian, &languageCode)
	if err != nil {
		return fmt.Errorf("failed to read language code: %v", err)
	}

	var localIP uint32
	err = binary.Read(reader, binary.LittleEndian, &localIP)
	if err != nil {
		return fmt.Errorf("failed to read local IP: %v", err)
	}
	state.ClientLocalIP = localIP

	var timezoneBias int32
	err = binary.Read(reader, binary.LittleEndian, &timezoneBias)
	if err != nil {
		return fmt.Errorf("failed to read time zone bias: %v", err)
	}
	state.TimezoneBias = timezoneBias

	var mpqLocaleId uint32
	err = binary.Read(reader, binary.LittleEndian, &mpqLocaleId)
	if err != nil {
		return fmt.Errorf("failed to read MPQ locale ID: %v", err)
	}

	var userLanguageId uint32
	err = binary.Read(reader, binary.LittleEndian, &userLanguageId)
	if err != nil {
		return fmt.Errorf("failed to read user language ID: %v", err)
	}
	state.LocaleUserLanguageId = userLanguageId

	state.CountryCode, err = ReadNullTerminatedByteArray(reader)
	if err != nil {
		return fmt.Errorf("failed to read country code: %v", err)
	}

	state.CountryName, err = ReadNullTerminatedByteArray(reader)
	if err != nil {
		return fmt.Errorf("failed to read country name: %v", err)
	}

	state.PingCookie = rand.Uint32()
	pingReply, err := WriteSID_PING(state.PingCookie)
	if err == nil {
		err = WriteSID(state.Conn, pingReply)
	}
	if err != nil {
		return fmt.Errorf("failed to write ping reply: %v", err)
	}

	return nil
}

func ReadNullTerminatedByteArray(r io.Reader) ([]byte, error) {
	var buf []byte
	var b [1]byte
	for {
		_, err := r.Read(b[:])
		if err != nil {
			return nil, err
		}
		if b[0] == 0 {
			break
		}
		buf = append(buf, b[0])
	}
	return buf, nil
}

func WriteSID(conn net.Conn, reply *message.Message) error {
	if reply.Length < 4 || reply.Length > 0xFFFF || reply.Length != uint16(4+len(reply.Body)) {
		return fmt.Errorf("invalid message reply length (expected 4-65535, got %d)", reply.Length)
	}

	buffer := make([]byte, reply.Length)
	buffer[0] = 0xFF
	buffer[1] = byte(reply.ID)
	buffer[2] = byte(reply.Length & 0xFF)
	buffer[3] = byte((reply.Length >> 8) & 0xFF)
	copy(buffer[4:], reply.Body)

	_, err := conn.Write(buffer)
	if err == nil {
		log.Printf("(%s) message ([0x%02X] %s) sent", conn.RemoteAddr(), reply.ID, message.MessageIdToName(reply.ID))
	}
	return err
}

func WriteSID_NULL() (*message.Message, error) {
	return &message.Message{
		ID:     message.SID_NULL,
		Length: 4,
		Body:   make([]byte, 0),
	}, nil
}

func WriteSID_PING(cookie uint32) (*message.Message, error) {
	buffer := &bytes.Buffer{}
	err := binary.Write(buffer, binary.LittleEndian, &cookie)
	if err != nil {
		return nil, err
	}

	return &message.Message{
		ID:     message.SID_PING,
		Length: uint16(4 + buffer.Len()),
		Body:   buffer.Bytes(),
	}, nil
}
