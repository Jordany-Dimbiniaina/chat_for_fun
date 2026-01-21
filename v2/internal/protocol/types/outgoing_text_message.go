package types

import "time"


type OutgoingTextMessage struct {
	Host    string
	Content string
	Sender  string
	Time time.Time
}

func (m *OutgoingTextMessage) GetHost() string {
	return m.Host
}
func (m *OutgoingTextMessage) GetContent() string {
	return m.Content
}
