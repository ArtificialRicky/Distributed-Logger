package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	grep "cs425.com/mp1/grepService"
)

func main() {
	grep_service := new(grep.GrepService)
	rpc.Register(grep_service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		fmt.Println("Listening...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		rpc.ServeConn(conn)
	}
}
