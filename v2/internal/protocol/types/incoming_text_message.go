package types

type IncomingTextMessage struct {
	Host    string
	Content string
}

func (m *IncomingTextMessage) GetHost() string {
	return m.Host
}
func (m *IncomingTextMessage) GetContent() string {
	return m.Content
}
