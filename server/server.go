package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/Jordany_dimbiniaina/chatForFun/message"
)


func handleNewConnection(conn net.Conn)  {

	out := make(chan message.Message, 10) // POURQUOI LE CHIFFRE ? BACKPRESURE C'EST QUI ??? 
	in := make(chan message.Message, 10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go message.OutgoingMessageHandler(ctx, out, conn)
	go message.IncomingMessageHadnler(ctx, in, cancel, conn, "||")

	out <- message.Message{
		Sender: "SERVER",
		Host: conn.RemoteAddr().String(),
		Content: fmt.Sprintf("WELCOME TO THE CHAT SERVER : %s", conn.RemoteAddr().String()),
	}
	
	for {
		select {
			case <- ctx.Done():
				log.Printf("[EVENT] Closed Connection : %s\n", conn.RemoteAddr().String())
				return
			case incomingMssage := <- in:
				out <- incomingMssage
		}
	}
}







type Server struct {
	Addr string
}

func (server Server) Start() (net.Listener, error) {
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return nil, err
	}
	
	log.Printf("Start server")
	fmt.Printf("TCP server running on : %s \n", server.Addr)
	return ln, nil
}

func (server Server) Serve(ln net.Listener)  {
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to connect : %v \n", err)
			continue
		}
		go handleNewConnection(conn)
	}

}