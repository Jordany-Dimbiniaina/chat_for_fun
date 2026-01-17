package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)



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


