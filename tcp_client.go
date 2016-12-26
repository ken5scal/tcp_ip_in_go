package main

import (
	"net"
	"fmt"
)

var TARGET_HOST = "www.google.co.jp"
var TARGET_PORT = "80"

func main() {
	// Create TCP Socket

	// Connect ot Server
	ip, err := net.ResolveIPAddr("ip", TARGET_HOST)
	if err != nil {
		panic(err)
	}
	fmt.Println(ip)


	// Senda data

	// Receive data
}