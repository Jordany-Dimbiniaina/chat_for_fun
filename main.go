package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)



const SERVER_ADDR = "127.0.0.1:2610"

func main() {

	
	logFile, err := os.OpenFile("log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v \n", err) 
    }
    defer logFile.Close()
    log.SetOutput(logFile)

	
	ln, err := net.Listen("tcp", SERVER_ADDR)
    if err != nil {
    	log.Fatalf("Failed to create the server : %v \n", err)
    }
	defer ln.Close()
	log.Printf("Start server")
	fmt.Printf("TCP server running on : %s \n", SERVER_ADDR)
    for {
    	conn, err := ln.Accept()
    	if err != nil {
    		fmt.Printf("Failed to connect : %v \n", err)
    	}
    	go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn)  {
	defer conn.Close()
	const GREETNGS = "WELCOME TO MY TCP CHANEL \n"
	if conn == nil {
		return
	}
	log.Printf("Connected : %s", conn.RemoteAddr())
	_ , err := conn.Write([]byte(GREETNGS))
	if err != nil {
		if err.Error() == "use of closed network connection" {
            log.Fatalf("Connection closed locally")
        }
		log.Printf("FAiled to write : %v \n", err)
		return
	}

	readBuffer := make([]byte, 5)
	for {
		n, err = conn.Read(readBuffer)
		if err == io.EOF {
			log.Printf("Connection closed by : %s \n", conn.RemoteAddr())
			return
		}
		if err != nil {
			if err.Error() == "use of closed network connection" {
				log.Fatalf("Connection closed locally")
			}
			log.Printf("Failed to read : %v \n", err)
			return
		}
		fmt.Printf(string(readBuffer[:n]))
	}
}

