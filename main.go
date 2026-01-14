package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const SERVER_ADDR = ":2610"

func main() {

	clients := make(map[string]net.Conn, 0)

	
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
    	go handleConnection(conn, clients)
    }
}

func handleConnection(conn net.Conn, clients map[string]net.Conn) {
	
	defer conn.Close()
	clients[conn.RemoteAddr().String()] = conn
	// fmt.Printf("CLIENTS : %v \n", clients)

	const GREETNGS = "WELCOME TO MY TCP CHANEL \n"
	const RECEIVED_HANDSHAKE = "RECEIVED \n"
	const NO_HOST_MENTIONED = "NO HOST MENTIONED OR INVALID HOST \n"
	const INVALID_HOST = "INVALID HOST \n"
	const MESSAGE_FORMAT_ERROR = "MESSAGE FORMAT ERROR \n"
	const UNREACHABLE_HOST = "UNREACHABLE HOST \n"
	


	if conn == nil {
		return
	}


	log.Printf("Connected : %s", conn.RemoteAddr())


	_ , err := conn.Write([]byte(fmt.Sprintf("%s %s\n", GREETNGS, conn.RemoteAddr())))

	if err != nil {
		if err.Error() == "use of closed network connection" {
            log.Fatalf("Connection closed locally")
        }
		log.Printf("Failed to write : %v \n", err)
		return
	}

	avalaibleUser := showAvalaibleUsers(clients)
	_ , err = conn.Write([]byte(fmt.Sprintf("%s \n", avalaibleUser)))
	if err != nil {
		if err.Error() == "use of closed network connection" {
            log.Fatalf("Connection closed locally")
        }
		log.Printf("Failed to write : %v \n", err)
		return
	}


	readBuffer := make([]byte, 1)
	var builder strings.Builder

	for {
		n, err := conn.Read(readBuffer)
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
		
		builder.WriteString(string(readBuffer[:n]))
		if string( readBuffer[:n]) == "*" {
			stream := builder.String()

			host, err := getHost(stream, "|")
			if err != nil {
				_ , err := conn.Write([]byte(MESSAGE_FORMAT_ERROR))
				if err != nil {
					if err.Error() == "use of closed network connection" {
						log.Fatalf("Connection closed locally")
					}
					log.Printf("Failed to write : %v \n", err)
					
				}
				return
			}

			// if validHost := validateHost(host); !validHost {
			// 	_ , err := conn.Write([]byte(INVALID_HOST))
			// 	if err != nil {
			// 		if err.Error() == "use of closed network connection" {
			// 			log.Fatalf("Connection closed locally")
			// 		}
			// 		log.Printf("Failed to write : %v \n", err)					
			// 	}
			// 	return
			// }

			hostConn , errConn := getConnByHost(host, clients)
			if errConn != nil && hostConn == nil {
				_ , err := conn.Write([]byte(UNREACHABLE_HOST))
				if err != nil {
					if err.Error() == "use of closed network connection" {
						log.Fatalf("Connection closed locally")
					}
					log.Printf("Failed to write : %v \n", err)					
				}
				return
			}

			_ , err = conn.Write([]byte(RECEIVED_HANDSHAKE))
			if err != nil {
				if err.Error() == "use of closed network connection" {
					log.Fatalf("Connection closed locally")
				}
				log.Printf("Failed to write : %v \n", err)		
				return			
			}
			messge := fmt.Sprintf("%s : %s \n", conn.RemoteAddr(), stream[len(host)+1:len(stream)-1])
			_ , err = hostConn.Write([]byte(messge))
			if err != nil {
				if err.Error() == "use of closed network connection" {
					log.Fatalf("Connection closed locally")
				}
				log.Printf("Failed to write : %v \n", err)		
				return			
			}

			fmt.Printf("%s -> %s : %s\n", conn.RemoteAddr(), host, stream[len(host)+1:len(stream)-1])
			
			builder.Reset()
			
		}	
	}
}

func getHost (stream , delimiter string) (string, error) {
	indexOfDelimiter, err := getIndexOfDelimiter(stream, delimiter); 
	if err != nil {
		return "", err
	}
	host := stream[:indexOfDelimiter]
	return host, nil
}

func validateHost(host string) bool {
	if len(host) < 13 || len(host) > 21 {
		return false
	}
	return true
}

func getIndexOfDelimiter (stream , delimiter string) (int, error) {
	indexOfDelimiter := strings.Index(stream, delimiter)
	if indexOfDelimiter == -1 {
		return -1, fmt.Errorf("Delimiter not found")
	}
	
	return indexOfDelimiter, nil
}

func getConnByHost(host string, clients map[string]net.Conn) (net.Conn, error) {
	conn, exists := clients[host]
	if !exists {
		return nil, fmt.Errorf("Host not found")
	}
	return conn, nil
}


func showAvalaibleUsers(clients map[string]net.Conn) string {
	var builder strings.Builder
	builder.WriteString("CONNECTED USERS \n")
	for host := range clients {
		builder.WriteString(fmt.Sprintf("\t %s \n", host))
	}
	return builder.String()
}