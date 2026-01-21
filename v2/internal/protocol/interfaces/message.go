package interfaces

import (
	"io"
)

type Message interface {
	GetHost() string
	GetContent() string
}

type MessageReader interface {
	ReadMessage(io.Reader) (Message, error)
}



