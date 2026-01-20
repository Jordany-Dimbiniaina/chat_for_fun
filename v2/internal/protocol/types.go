package protocol

import (
	"bufio"
	"io"
	"strings"
)


type TextMessageReader struct {}

func (reader *TextMessageReader) ReadMessage(r io.Reader, delimiter string) (Message, error) {
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
		host:    host,
		content:contentBuilder.String(),
	}, scanner.Err()
}