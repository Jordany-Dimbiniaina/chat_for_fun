package types

import (
	"fmt"
	"io"
	"github.com/Jordany-Dimbiniaina/chat_for_fun/v2/internal/protocol/interfaces"
)


type TextMessageWriter  struct {
	writer io.Writer
}

func (messageWriter *TextMessageWriter) WriteMessage(message interfaces.Message ) (int, error) {
	content := fmt.Sprintf("[%s] (%s) : %s\n", message.GetSender(), message.GetTime(), message.GetContent())
	n, err := messageWriter.writer.Write([]byte(content))
	return n, err
}

func NewTextMessageWriter( writer io.Writer) interfaces.MessageWriter  {
	return &TextMessageWriter{
		writer: writer,
	}
}