package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const GREETNGS = "WELCOME TO MY TCP CHANEL \n"
const RECEIVED_HANDSHAKE = "RECEIVED \n"
const NO_HOST_MENTIONED = "NO HOST MENTIONED OR INVALID HOST \n"
const INVALID_HOST = "INVALID HOST \n"
const MESSAGE_FORMAT_ERROR = "MESSAGE FORMAT ERROR \n"
const UNREACHABLE_HOST = "UNREACHABLE HOST \n"



func ValidateHost(host string) bool {
	if len(host) < 13 || len(host) > 21 {
		return false
	}
	return true
}

func GetIndexOfDelimiter(stream, delimiter string) (int, error) {
	indexOfDelimiter := strings.Index(stream, delimiter)
	if indexOfDelimiter == -1 {
		return -1, fmt.Errorf("Delimiter not found")
	}

	return indexOfDelimiter, nil
}

func GetConnByHost(host string, clients map[string]net.Conn) (net.Conn, error) {
	conn, exists := clients[host]
	if !exists {
		return nil, fmt.Errorf("Host not found")
	}
	return conn, nil
}

func AvalaibleUsers(clients map[string]net.Conn) string {
	var builder strings.Builder
	builder.WriteString("CONNECTED USERS \n")
	for host := range clients {
		builder.WriteString(fmt.Sprintf("\t %s \n", host))
	}
	return builder.String()
}


func ConfigureLog()  {
	logFile, err := os.OpenFile("log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v \n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
}
