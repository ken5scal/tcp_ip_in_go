package main

import (
	"net"
	"fmt"
	"os"
	"github.com/labstack/gommon/log"
)

var BIND_IP = "127.0.0.1"
var BIND_PORT = "8080"

func main() {
	//tcpAddr, err := net.ResolveTCPAddr("tcp", BIND_IP + ":" + BIND_PORT)
	//checkServerError(err)

	ln, err := net.Listen("tcp", BIND_IP + ":" + BIND_PORT)
	checkServerError(err)
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
	//daytime := time.Now().String()
	//c.Write([]byte(daytime))

	var buf [512]byte
	for {
		n, err := c.Read(buf[0:])
		fmt.Printf("n: %v\n", n)
		if err != nil {
			return
		}
		_, err2 := c.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkServerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}