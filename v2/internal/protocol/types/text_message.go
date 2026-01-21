package types

import "time"

type TextMessage struct {
	Host    string
	Content string
	Sender  string
	Time    time.Time
}

func (m TextMessage) GetHost() string {
	return m.Host
}
func (m TextMessage) GetContent() string {
	return m.Content
}

func (m TextMessage) GetTime() time.Time {
	return m.Time
}

func (m TextMessage) GetSender() string {
	return m.Sender
}