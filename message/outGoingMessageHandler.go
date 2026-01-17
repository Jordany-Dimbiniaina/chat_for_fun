package message

import (
	"context"
	"fmt"
	"net"
)

func OutgoingMessageHandler(ctx context.Context, out chan Message, conn net.Conn) {

	
	for {
		select {
		case <-ctx.Done():
			return
		case message := <-out:
			fmt.Fprintf(conn, "%s -> %s : %s \n", message.Sender, message.Host, message.Content)
		}
	}
}