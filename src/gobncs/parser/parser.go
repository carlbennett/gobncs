package gobncs

type MessageId byte

const (
	SID_NULL MessageId = 0x00
	SID_PING MessageId = 0x25
)

func ReceiveMessage(buffer []byte) {

}

func WriteMessage(messageId uint8, messageData []byte) []byte {
	messageLength := len(messageData) + 4
	buffer := make([]byte, messageLength)

	buffer[0] = 0xFF
	buffer[1] = messageId
	buffer[2] = byte(messageLength & 0xFF)
	buffer[3] = byte(messageLength >> 8)

	copy(buffer[4:], messageData)

	return buffer
}
