package message

import (
	"context"
	"fmt"
	"net"
	"github.com/Jordany_dimbiniaina/chatForFun/interfaces"
	"github.com/Jordany_dimbiniaina/chatForFun/utils"
)


func OutgoingMessageHandler(ctx context.Context, out chan Message, sender net.Conn, clientStore interfaces.ClientStore)  {
	for {
		select {
		case <-ctx.Done():
			return
		case message := <-out:
			hostConn := utils.GetHostConn(message.Host, clientStore)			
			if hostConn == nil {
				message.Sender = "SERVER"
				message.SystemMessage = true
				message.Content = "UNREACHABLE HOST \n"
				hostConn = sender
			} 
			if message.SystemMessage {
				fmt.Fprintf(hostConn, " %s \n",message.Content)
			} else {
				fmt.Fprintf(hostConn, "%s : %s \n", message.Sender, message.Content)
			}
		}
	}
}

