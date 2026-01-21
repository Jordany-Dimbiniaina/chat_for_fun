package types

import (
	"fmt"
	"io"
)


type TextMessageWriter  struct {}

func (messageWriter *TextMessageWriter) WriteMessage(writer io.Writer, message OutgoingTextMessage) (int, error) {
	content := fmt.Sprintf("[%s] (%s) : %s\n", message.Sender, message.Time, message.Content)
	n, err := writer.Write([]byte(content))
	return n, err
}