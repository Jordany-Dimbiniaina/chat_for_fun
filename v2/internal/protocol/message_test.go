package protocol

import (
	"strings"
	"testing"
)



func TestTextMessageReader(t *testing.T) {

	stream := "192.168.1.1:26100\nHello My friend!\nHow are you?\n||"
	reader := strings.NewReader(stream)
	textMessageReader := &TextMessageReader{}
	delimiter := "||"

	message, err := textMessageReader.ReadMessage(reader, delimiter)
	sameHost := message.Host() == "192.168.1.1:26100"
	sameContent := message.Content() == "Hello My friend!\nHow are you?\n"

	if !sameContent {
		t.Errorf("Expected content %s, got %s, err  : %v", "Hello My friend!\nHow are you?\n", message.Content(), err)
	}

	if !sameHost {
		t.Errorf("Expected host %s, got %s, err : %v", "192.168.1.1:26100", message.Host(), err )
	}

}

