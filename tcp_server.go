package main

import (
	"net"
	"fmt"
	"os"
	"time"
)

var BIND_IP = "127.0.0.1"
var BIND_PORT = "8080"

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", BIND_IP + ":" + BIND_PORT)
	checkServerError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	daytime := time.Now().String()
	c.Write([]byte(daytime))
}

func checkServerError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v", err.Error())
		os.Exit(1)
	}
}