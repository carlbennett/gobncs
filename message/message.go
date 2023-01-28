package message

import (
	"encoding/binary"
	"fmt"
	"io"
)

type Message struct {
	ID     byte
	Length uint16
	Body   []byte
}

func ReadMessage(conn io.Reader) (*Message, error) {
	var header [4]byte
	_, err := io.ReadFull(conn, header[:])
	if err != nil {
		return nil, err
	}

	if header[0] != 0xFF {
		return nil, fmt.Errorf("invalid message header")
	}

	length := binary.LittleEndian.Uint16(header[2:4])
	body := make([]byte, length-4)
	_, err = io.ReadFull(conn, body)
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:     header[1],
		Length: length,
		Body:   body,
	}, nil
}

func ReadProtocolByte(conn io.Reader) (byte, error) {
	var buf [1]byte
	_, err := io.ReadFull(conn, buf[:])
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}
