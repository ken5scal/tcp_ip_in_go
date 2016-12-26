package main

import (
	"net"
	"fmt"
	"os"
)

var TARGET_HOST = "www.google.co.jp"
var TARGET_PORT = "80"

func main() {
	// Create TCP Socket

	/**
	Connect ot Server
	 */
	// Resolve IP addr
	ip, err := net.ResolveTCPAddr("tcp", TARGET_HOST + ":" + TARGET_PORT)
	checkError(err)

	fmt.Println(ip)

	// Senda data

	// Receive data
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v", err.Error())
		os.Exit(1)
	}
}