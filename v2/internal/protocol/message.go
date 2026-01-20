package protocol

import (
	"io"
)



func NewIncomingTextMessage(reader io.Reader, messageReader MessageReader) (Message, error) {
	message, err := messageReader.ReadMessage(reader)
	return message, err
}