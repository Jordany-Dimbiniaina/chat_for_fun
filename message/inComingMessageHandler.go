package message

import (
	"bufio"
	"context"
	"net"
	"strings"
)


func IncomingMessageHadnler(ctx context.Context, in chan Message, cancel context.CancelFunc , conn net.Conn, delimiter string) {

	scanner := bufio.NewScanner(conn)
	firstLine := true
	host := conn.RemoteAddr().String()
	var messageBuilder strings.Builder 
	defer messageBuilder.Reset()

    for scanner.Scan() {

		if firstLine {
			host = scanner.Text()
			firstLine = false
			continue
		}

		line := scanner.Text()
		if line != delimiter { // not the end of the message
			messageBuilder.WriteString(line)
			messageBuilder.WriteString("\n")
			continue
		}

		messageContent := strings.TrimRight(messageBuilder.String(), "\n")
        messageBuilder.Reset()

        if messageContent == "" {
            continue
        }
			
		
        select {
			case <-ctx.Done():
				return
			case in <- Message{
				Host: host,
				Sender: conn.RemoteAddr().String(),
				Content: messageContent,
			}:
        }
    }


	cancel()
}