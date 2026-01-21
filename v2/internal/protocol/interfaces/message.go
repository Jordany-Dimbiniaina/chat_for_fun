package interfaces

import (
	"time"
)

type Message interface {
	GetHost() string
	GetContent() string
	GetTime() time.Time 
	GetSender() string
}

type MessageReader interface {
	ReadMessage() (Message, error)
}

type MessageWriter interface {
	WriteMessage(Message) (int, error)
}



