package errors

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
)


func handleClientDeconnected(ctx context.Context,cancel context.CancelFunc, conn net.Conn) {
	log.Printf("Connection closed by : %s \n", conn.RemoteAddr())
	cancel() 
	return
}



func HandleServerEroor(ctx context.Context, cancel context.CancelFunc, err error)  {
	_, ok := err.(net.Error);
	switch  {
		case  errors.Is(err, net.ErrClosed) :
			log.Printf("[ERROR] Server closed listener")
		case ok :
			log.Printf("[ERROR] Temporary server error or timeout: %v", err)
		default :
			log.Printf("[ERROR] SERVER ERROR: %v", err)
	}
	
}

func HandleConnError(ctx context.Context, cancel context.CancelFunc, conn net.Conn, err error) {
    if err == nil {
        return
    }
    switch {
		case errors.Is(err, io.EOF):
			handleClientDeconnected(ctx, cancel, conn)
		default:
			log.Printf("[ERROR] : %v\n", err)
    }
    conn.Close()
}

