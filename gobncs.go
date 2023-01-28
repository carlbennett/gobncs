package main

import (
	"log"
	"net"
	"os"

	"github.com/carlbennett/gobncs/server"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("[gobncs] ")

	ln, _ := net.Listen("tcp", ":6112")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		go server.HandleConnection(conn)
	}
}
