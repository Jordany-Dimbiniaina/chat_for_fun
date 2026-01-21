package protocol

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/Jordany-Dimbiniaina/chat_for_fun/v2/internal/protocol/types"
)


func Test(t *testing.T) {
	var writer bytes.Buffer 
	message := types.OutgoingTextMessage {
		Host : "192.168.1.1:2610",
		Sender : "192.168.1.1:2160",
		Time : time.Now(),
		Content : "Hello My Friend",
	}
	textMessageWriter := types.TextMessageWriter{}
	wantedContent := fmt.Sprintf("[%s] (%s) : %s\n", message.Sender, message.Time.String(), message.Content)
	_, err := textMessageWriter.WriteMessage(&writer, message)
	if err != nil &&  writer.String() != wantedContent {
		t.Errorf("Failed on writting message, got : %s, wanted : %s, ERROR : %v", writer.String(), message.Content, err )
	}
}