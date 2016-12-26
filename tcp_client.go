package main

import (
	"net"
	"fmt"
	"os"
	"io/ioutil"
)

var TARGET_HOST = "www.google.co.jp"
var TARGET_PORT = "80"

func main() {
	// Create TCP Socket

	/**
	Connect ot Server
	 */
	// Resolve IP addr
	tcpAddr, err := net.ResolveTCPAddr("tcp", TARGET_HOST + ":" + TARGET_PORT)
	checkError(err)

	fmt.Println(tcpAddr.IP)
	fmt.Println(tcpAddr.Port)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	//conn.SetWriteDeadline(time.Now().Add(time.Minute))
	// Senda data
	//status, err := conn.Write([]byte("GET  HTTP/1.1\r\nHost: google.com\r\n\r\n"))
	status, err := conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	checkError(err)
	fmt.Printf("status code: %v\n", string(status))

	// Receive data
	result_byte, err := ioutil.ReadAll(conn)
	//result_byte := make([]byte, 1024)
	//_, err = conn.Read(result_byte)
	checkError(err)
	//fmt.Printf("response code: %v\n", string(result))
	fmt.Printf("response code: %v\n", string(result_byte))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %v", err.Error())
		os.Exit(1)
	}
}