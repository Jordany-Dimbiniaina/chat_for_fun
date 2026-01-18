package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/Jordany_dimbiniaina/chatForFun/interfaces"
	"github.com/Jordany_dimbiniaina/chatForFun/message"
)

func handleNewConnection(conn net.Conn, clientStore interfaces.ClientStore) {

	out := make(chan message.Message, 10) // POURQUOI LE CHIFFRE ? BACKPRESURE C'EST QUI ???
	in := make(chan message.Message, 10)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go message.OutgoingMessageHandler(ctx, out, conn, clientStore)
	go message.IncomingMessageHandler(ctx, in, cancel, conn, "||")

	out <- GreetingsMessage(conn, clientStore)
	// out <- message.Message{
	// 	Sender: "SERVER",
	// 	SystemMessage: true,
	// 	Content: fmt.Sprintf("CONNECTED USERS : \n  %s ", clientStore.List()),
	// 	Host : conn.RemoteAddr().String(),
	// }

	for {
		select {
		case <-ctx.Done():
			log.Printf("[EVENT] Closed Connection : %s\n", conn.RemoteAddr().String())
			return
		case incomingMssage := <-in:
			out <- incomingMssage
		}
	}
}

type Server struct {
	Addr        string
	ClientStore interfaces.ClientStore
}

func (server Server) Start() (net.Listener, error) {
	ln, err := net.Listen("tcp", "")
	if err != nil {
		return nil, err
	}
	log.Printf("Start server")
	fmt.Printf("TCP server running on : %s \n", ln.Addr().String())
	return ln, nil
}

func (server Server) Serve(ln net.Listener) {
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to connect : %v \n", err)
			continue
		}
		server.storeNewUser(conn)
		fmt.Printf("%v \n", server.ClientStore)
		go handleNewConnection(conn, server.ClientStore)
	}
}

func (server Server) storeNewUser(conn net.Conn) {
	server.ClientStore.Store(conn.RemoteAddr().String(), conn)
}

func NewServer(addr string, clientStore interfaces.ClientStore) *Server {
	return &Server{
		Addr:        addr,
		ClientStore: clientStore,
	}
}

func GreetingsMessage(conn net.Conn, clientStore interfaces.ClientStore) message.Message {
	return message.Message{
		Sender:  "SERVER",
		SystemMessage: true,
		Host:    conn.RemoteAddr().String(),
		Content: fmt.Sprintf("WELCOME : %s \n", conn.RemoteAddr().String()),
	}
}



