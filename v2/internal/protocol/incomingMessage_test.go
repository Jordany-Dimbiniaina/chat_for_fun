package protocol

import (
	"strings"
	"testing"
	"github.com/Jordany-Dimbiniaina/chat_for_fun/v2/internal/protocol/types"
)


func TestTextMessageReader(t *testing.T) {

	stream := "192.168.1.1:26100\nHello My friend!\nHow are you?\n||"
	reader := strings.NewReader(stream)
	delimiter := "||"

	textMessageReader := types.NewTextMessageReader(reader, delimiter)
	message, err := textMessageReader.ReadMessage()
	
	sameHost := message.GetHost() == "192.168.1.1:26100"
	sameContent := message.GetContent() == "Hello My friend!\nHow are you?\n"

	if !sameContent {
		t.Errorf("Expected content %s, got %s, err  : %v", "Hello My friend!\nHow are you?\n", message.GetContent(), err)
	}

	if !sameHost {
		t.Errorf("Expected host %s, got %s, err : %v", "192.168.1.1:26100", message.GetHost(), err )
	}

}

