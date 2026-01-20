package types

import (
	"bufio"
	"io"
	"strings"

	protocol "github.com/Jordany-Dimbiniaina/chat_for_fun/v2/internal/protocol/interfaces"
)


type TextMessageReader struct {}

func (reader *TextMessageReader) ReadMessage(r io.Reader, delimiter string) (protocol.Message, error) {
	scanner := bufio.NewScanner(r)
	firstLine := true
	host := ""
	contentBuilder := strings.Builder{}

	for scanner.Scan() {
		if firstLine {
			firstLine = false
			host = scanner.Text()
			continue
		}
		line := scanner.Text()
		if line == delimiter {
			break
		}
		contentBuilder.WriteString(line)
		contentBuilder.WriteString("\n")
	}
	
	return &IncomingTextMessage{
		Host:    host,
		Content:contentBuilder.String(),
	}, scanner.Err()
}