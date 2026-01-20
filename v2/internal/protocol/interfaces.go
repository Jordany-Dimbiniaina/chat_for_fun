package protocol

import (
	"io"
)

type Message interface {
	Host() string
	Content() string
}

type IncomingTextMessage struct {
	host string
	content string
}

func (m *IncomingTextMessage) Host() string {
	return m.host
}		
func (m *IncomingTextMessage) Content() string {
	return m.content
}

type MessageReader interface {
	ReadMessage(io.Reader) (Message, error)
}
