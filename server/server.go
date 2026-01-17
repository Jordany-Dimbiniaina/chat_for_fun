package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"github.com/Jordany_dimbiniaina/chatForFun/message"
)

func handleNewConnection(conn net.Conn, clientStore ClientStore) {

	out := make(chan message.Message, 10) // POURQUOI LE CHIFFRE ? BACKPRESURE C'EST QUI ???
	in := make(chan message.Message, 10)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go message.OutgoingMessageHandler(ctx, out, conn)
	go message.IncomingMessageHandler(ctx, in, cancel, conn, "||")

	out <- GreetingsMessage(conn, clientStore)

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
	ClientStore ClientStore
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

func (server Server) Serve(ln net.Listener) {
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to connect : %v \n", err)
			continue
		}
		server.storeNewUser(conn)
		go handleNewConnection(conn, server.ClientStore)
	}
}

func (server Server) storeNewUser(conn net.Conn) {
	server.ClientStore.Store(conn.RemoteAddr().String(), conn)
}

func NewServer(addr string, clientStore ClientStore) *Server {
	return &Server{
		Addr:        addr,
		ClientStore: clientStore,
	}
}

func GreetingsMessage(conn net.Conn, clientStore ClientStore) message.Message {
	
	return message.Message{
		Sender:  "SERVER",
		Host:    conn.RemoteAddr().String(),
		Content: fmt.Sprintf("WELCOME : %s \n", conn.RemoteAddr().String()),
	}
}

