package types

import (
	"bufio"
	"io"
	"strings"
	"github.com/Jordany-Dimbiniaina/chat_for_fun/v2/internal/protocol/interfaces"
)

type TextMessageReader struct{
	reader io.Reader
	delimiter string
}

func (reader *TextMessageReader) ReadMessage() (interfaces.Message, error) {
	scanner := bufio.NewScanner(reader.reader)
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
		if line == reader.delimiter {
			break
		}
		contentBuilder.WriteString(line)
		contentBuilder.WriteString("\n")
	}

	return &TextMessage{
		Host:    host,
		Content: contentBuilder.String(),
	}, scanner.Err()
}


func NewTextMessageReader(reader io.Reader, delimiter string) interfaces.MessageReader {
	return &TextMessageReader{
		reader: reader,
		delimiter : delimiter,
	}
}