package main

import (
	"net"
	"fmt"
	"os"
)

var BIND_IP = "127.0.0.1"
var BIND_PORT = ":8080"

func main() {
	ln, err := net.Listen("tcp", BIND_IP+BIND_PORT)
	checkServerError(err)
	for {
		conn, err := ln.Accept()
		checkServerError(err)
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {

}

func checkServerError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v", err.Error())
		os.Exit(1)
	}
}