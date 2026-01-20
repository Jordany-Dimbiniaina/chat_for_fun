package protocol

import (
	"io"
	"github.com/Jordany-Dimbiniaina/chat_for_fun/v2/internal/protocol/interfaces"
)



func NewIncomingTextMessage(reader io.Reader, messageReader interfaces.MessageReader) (interfaces.Message, error) {
	message, err := messageReader.ReadMessage(reader)
	return message, err
}

