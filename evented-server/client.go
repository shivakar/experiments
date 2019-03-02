package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	hostname := "127.0.0.1"
	port := "8181"
	if len(os.Args) < 3 {
		fmt.Println("Usage: client HOSTNAME PORT")
		fmt.Println("Using defaults:")
	} else {
		fmt.Println("Using the following values:")
		hostname = os.Args[1]
		port = os.Args[2]
	}
	fmt.Println("\tHOSTNAME : ", hostname)
	fmt.Println("\tPORT     : ", port)

	serverAddr := fmt.Sprintf("%s:%s", hostname, port)
	// Checking host address
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		panic(err)
	}
	// Connect to the server
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

}
