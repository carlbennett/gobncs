package main

import (
	"log"
	"net"
	"os"

	"github.com/carlbennett/gobncs/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("[gobncs] ")

	ln, _ := net.Listen("tcp", ":6113")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		go server.HandleConnection(conn)
	}
}
