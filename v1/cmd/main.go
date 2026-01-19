package main

import (
	"log"
	"github.com/Jordany_dimbiniaina/chatForFun/server"
	"github.com/Jordany_dimbiniaina/chatForFun/types"
	"github.com/Jordany_dimbiniaina/chatForFun/utils"
)

const SERVER_ADDR = "0.0.0.0:2610"

func main ()  {
	logfile := utils.ConfigureLog()
	defer logfile.Close()
	clientStore := types.NewTCPClientStore()
	server := server.NewServer(SERVER_ADDR, clientStore)

	ln, err := server.Start()
	if err != nil {
		log.Fatalf("Failed to start server : %v \n", err)
	} 
	log.Printf("[EVENT] Starting server")
	server.Serve(ln)
}