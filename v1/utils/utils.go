package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"github.com/Jordany_dimbiniaina/chatForFun/interfaces"
)


func GetHostConn(addr string, clientStore interfaces.ClientStore)  interfaces.ReadWriteCloser {
	hostConn, connected := clientStore.Load(addr)
	if !connected {
		return nil
	}
	return hostConn
}


func AvalaibleUsers(clients map[string]net.Conn) string {
	var builder strings.Builder
	builder.WriteString("CONNECTED USERS \n")
	for host := range clients {
		builder.WriteString(fmt.Sprintf("\t %s \n", host))
	}
	return builder.String()
}


func ConfigureLog()  *os.File{
	logFile, err := os.OpenFile("log/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v \n", err)
	}
	log.SetOutput(logFile)
	return logFile
}


