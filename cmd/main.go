package main

import (
	"log"
	"github.com/Jordany_dimbiniaina/chatForFun/server"
	"github.com/Jordany_dimbiniaina/chatForFun/utils"
)


const SERVER_ADDR = "localhost:2610"

func main ()  {
	
	utils.ConfigureLog()
	server := server.Server{
		Addr: SERVER_ADDR,
	}
	ln, err := server.Start()
	if err != nil {
		log.Fatalf("Failed to start server : %v \n", err)
	}
	server.Serve(ln)
}